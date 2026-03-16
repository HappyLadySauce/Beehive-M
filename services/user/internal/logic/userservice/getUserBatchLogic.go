package userservicelogic

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/HappyLadySauce/Beehive-M/pkg/code"
	"github.com/HappyLadySauce/Beehive-M/services/user/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/user/pb"

	"github.com/HappyLadySauce/errors"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type GetUserBatchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserBatchLogic {
	return &GetUserBatchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量按 user_id 查（会话列表、群成员列表等）
func (l *GetUserBatchLogic) GetUserBatch(in *pb.GetUserBatchRequest) (*pb.GetUserBatchResponse, error) {
	// 1.1 参数校验: user_ids 不能为空
	userIds := in.GetUserIds()
	if len(userIds) == 0 {
		// 1.2 如果 user_ids 为空，则返回参数错误
		l.Logger.Errorf("user_ids is required")
		return nil, errors.WithCode(code.CodeInvalidParam, "user_ids is required")
	}

	// 1.3 防御性限制: 单次最多允许查询 500 个 user_id
	if len(userIds) > 500 {
		l.Logger.Errorf("too many user_ids, max allowed is 500, got: %d", len(userIds))
		return nil, errors.WithCode(code.CodeInvalidParam, "too many user_ids, max allowed is 500")
	}

	// 2.1 去重，避免重复查询
	userIds = lo.Uniq(userIds)

	// 3.1 构建 key 列表: user:profile:{id}
	keys := make([]string, len(userIds))
	for i, userId := range userIds {
		keys[i] = fmt.Sprintf("user:profile:%d", userId)
	}

	// 3.2 user_id 到用户信息的映射，方便保持入参顺序
	userMap := make(map[int64]*pb.User, len(userIds))
	var missIds []int64

	// 3.3 查询缓存（MGet 对不存在的 key 返回 nil，不会返回 redis.Nil 错误）
	values, err := l.svcCtx.Redis.MGet(l.ctx, keys...).Result()
	if err != nil {
		// 3.4 如果 Redis 返回错误（包括连接异常等），记日志并全部回退到数据库查询
		l.Logger.Errorf("get user profile from redis failed, fallback to db only, err: %v", err)
		values = nil
		missIds = append(missIds, userIds...)
	}

	// 4.1 解析缓存命中的用户，记录遗漏的 user_id
	for i, val := range values {
		id := userIds[i]
		if val == nil {
			// 4.2 未命中缓存
			missIds = append(missIds, id)
			continue
		}

		strVal, ok := val.(string)
		if !ok || strVal == "" {
			// 4.3 类型异常或空字符串，当作未命中
			missIds = append(missIds, id)
			continue
		}

		var u pb.User
		if err := json.Unmarshal([]byte(strVal), &u); err != nil {
			// 4.4 与单查逻辑保持一致，解析失败直接报错
			l.Logger.Errorf("unmarshal user profile from redis failed: %v", err)
			return nil, errors.WithCode(code.CodeUnmarshalFailed, "unmarshal failed")
		}
		userMap[id] = &u
	}

	// 5.1 如果所有用户都在缓存中，则按入参顺序返回
	if len(missIds) == 0 {
		users := make([]*pb.User, 0, len(userIds))
		for _, id := range userIds {
			if u, ok := userMap[id]; ok {
				users = append(users, u)
			}
		}
		return &pb.GetUserBatchResponse{Users: users}, nil
	}

	// 6.1 查询数据库中遗漏的用户
	var dbUsers []*pb.User
	tx := l.svcCtx.DB.Where("user_id IN ?", missIds).Find(&dbUsers)
	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		// 6.2 所有遗漏的用户都不存在，继续返回已命中的部分
		l.Logger.Infof("some users not found in database, user_ids: %v", missIds)
	} else if tx.Error != nil {
		// 6.3 如果查询失败，并且不是记录不存在，则返回数据库错误
		l.Logger.Errorf("query users in database failed: %v", tx.Error)
		return nil, errors.WithCode(code.CodeDBQueryFailed, "query users in database failed")
	}

	// 7.1 将数据库中查到的用户写入 map，并回写缓存
	for _, u := range dbUsers {
		// 7.2 如果用户为 nil，则跳过
		if u == nil {
			continue
		}
		// 7.3 如果 user_id 为 0，则跳过
		id := u.GetUserId()
		if id == 0 {
			continue
		}

		userMap[id] = u

		// 7.4 序列化用户
		userBytes, err := json.Marshal(u)
		if err != nil {
			// 7.5 如果序列化失败，则返回序列化失败
			l.Logger.Errorf("marshal user profile to json failed: %v", err)
			return nil, errors.WithCode(code.CodeUnmarshalFailed, "marshal failed")
		}

		// 7.6 回写缓存
		// key: user:profile:{id}
		// value: user json
		err = l.svcCtx.Redis.Set(l.ctx, fmt.Sprintf("user:profile:%d", id), userBytes, 0).Err()
		if err != nil {
			// 7.7 如果回写缓存失败，则返回缓存错误
			l.Logger.Errorf("set user profile to redis failed: %v", err)
			return nil, errors.WithCode(code.CodeCacheSetFailed, "set user profile to cache failed")
		}
	}

	// 8.1 按照入参 user_ids 的顺序组装返回结果；未找到的用户直接跳过
	users := make([]*pb.User, 0, len(userIds))
	for _, id := range userIds {
		if u, ok := userMap[id]; ok {
			users = append(users, u)
		}
	}

	// 9. 返回结果
	return &pb.GetUserBatchResponse{Users: users}, nil
}
