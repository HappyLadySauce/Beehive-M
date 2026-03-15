package userservicelogic

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/HappyLadySauce/Beehive-M/services/user/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/user/pb"
	"github.com/HappyLadySauce/Beehive-M/pkg/code"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/HappyLadySauce/errors"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 按 user_id 查单个用户（会话头、资料页等）
func (l *GetUserLogic) GetUser(in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// 1.1 参数校验: user_id 不能为空
	id := in.GetUserId()
	if id == 0 {
		// 1.2 如果 user_id 为空，则返回参数错误
		l.Logger.Errorf("user_id is required")
		return nil, errors.WithCode(code.CodeInvalidParam, "user_id is required")
	}

	var user pb.User

	// 2.1 先查询缓存中是否存在用户
	// key: user:profile:{id}
	// value: user json
	userBytes, err := l.svcCtx.Redis.Get(l.ctx, fmt.Sprintf("user:profile:%d", id)).Bytes()
	// 2.2 如果缓存命中，则直接返回
	if err == nil && len(userBytes) > 0 {
		if err := json.Unmarshal(userBytes, &user); err == nil {
			return &pb.GetUserResponse{User: &user}, nil
		} else {
			// 2.3 如果解析失败，则返回解析失败
			l.Logger.Errorf("unmarshal user profile from redis failed: %v", err)
			return nil, errors.WithCode(code.CodeUnmarshalFailed, "unmarshal failed")
		}
	}

	// 3.1 如果缓存未命中，则查询数据库
	err = l.svcCtx.DB.Where("user_id = ?", id).First(&user).Error
	if err != nil {
		// 3.2 如果查询失败，并且是记录不存在，则返回用户不存在
		if errors.Is(err, gorm.ErrRecordNotFound) {
			l.Logger.Errorf("user not found in database: %v", err)
			return nil, errors.WithCode(code.CodeUserNotFound, "user not found in database")
		}
		// 3.3 如果查询失败，并且不是记录不存在，则返回数据库错误
		l.Logger.Errorf("query user in database failed: %v", err)
		return nil, errors.WithCode(code.CodeDBQueryFailed, "query user in database failed")
	}

	// 4. 回写缓存（传指针避免复制含 sync.Mutex 的 pb.User）
	userBytes, err = json.Marshal(&user)
	if err != nil {
		// 4.1 如果序列化失败，则返回序列化失败
		l.Logger.Errorf("marshal user profile to json failed: %v", err)
		return nil, errors.WithCode(code.CodeUnmarshalFailed, "marshal failed")
	}
	// 4.2 如果序列化成功，则回写缓存
	// key: user:profile:{id}
	// value: user json
	err = l.svcCtx.Redis.Set(l.ctx, fmt.Sprintf("user:profile:%d", id), userBytes, 0).Err()
	if err != nil {
		// 4.3 如果回写缓存失败，则返回缓存错误
		l.Logger.Errorf("set user profile to redis failed: %v", err)
		return nil, errors.WithCode(code.CodeCacheSetFailed, "set user profile to cache failed")
	}

	// 5. 返回用户信息
	return &pb.GetUserResponse{User: &user}, nil
}
