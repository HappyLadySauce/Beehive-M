package userservicelogic

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/HappyLadySauce/Beehive-M/pkg/code"
	"github.com/HappyLadySauce/Beehive-M/services/user/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/user/pb"

	"github.com/HappyLadySauce/errors"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type GetUserByAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByAccountLogic {
	return &GetUserByAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 按账号查（登录、搜索用）
func (l *GetUserByAccountLogic) GetUserByAccount(in *pb.GetUserByAccountRequest) (*pb.GetUserByAccountResponse, error) {
	// 1.1 参数校验: user_id 不能为空
	account := in.GetAccount()
	if account == "" {
		// 1.2 如果 account 为空，则返回参数错误
		l.Logger.Errorf("account is required")
		return nil, errors.WithCode(code.CodeInvalidParam, "account is required")
	}

	var user pb.User

	// 2.1 根据账号查询用户ID
	userId, err := l.getUserIdByAccount(account)
	if err != nil {
		return nil, err
	}

	// 3.1 先查询缓存中是否存在用户
	// key: user:profile:{id}
	// value: user json
	userBytes, err := l.svcCtx.Redis.Get(l.ctx, fmt.Sprintf("user:profile:%d", userId)).Bytes()
	// 3.2 如果 Redis 返回缓存命中，则直接返回
	if err == nil && len(userBytes) > 0 {
		// 3.3 如果缓存命中，则直接返回
		if err := json.Unmarshal(userBytes, &user); err == nil {
			return &pb.GetUserByAccountResponse{User: &user}, nil
		} else {
			// 3.4 如果解析失败，则返回解析失败
			l.Logger.Errorf("unmarshal user profile from redis failed: %v", err)
			return nil, errors.WithCode(code.CodeUnmarshalFailed, "unmarshal failed")
		}
	} else if errors.Is(err, redis.Nil) {
		// 3.5 如果 Redis 返回 key 不存在，记日志，继续走数据库
		l.Logger.Infof("get user profile from redis not found: %v", err)
	} else {
		// 3.6 如果 Redis 返回其他错误，记日志，继续走数据库
		l.Logger.Errorf("get user profile from redis failed: %v", err)
	}

	// 4.1 如果缓存未命中，则查询数据库
	tx := l.svcCtx.DB.Where("user_id = ?", userId).First(&user)
	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		// 4.2 如果查询失败，并且是记录不存在，则返回用户不存在
		l.Logger.Errorf("user not found in database: %v", tx.Error)
		return nil, errors.WithCode(code.CodeUserNotFound, "user not found in database")
	} else if tx.Error != nil {
		// 4.3 如果查询失败，并且不是记录不存在，则返回数据库错误
		l.Logger.Errorf("query user in database failed: %v", tx.Error)
		return nil, errors.WithCode(code.CodeDBQueryFailed, "query user in database failed")
	}

	// 5. 回写缓存
	userBytes, err = json.Marshal(&user)
	if err != nil {
		// 5.1 如果序列化失败，则记录日志
		l.Logger.Errorf("marshal user profile to json failed: %v", err)
	}
	// 5.2 如果序列化成功，则回写缓存
	// key: user:profile:{id}
	// value: user json
	err = l.svcCtx.Redis.Set(l.ctx, fmt.Sprintf("user:profile:%d", userId), userBytes, 0).Err()
	if err != nil {
		// 5.3 如果回写缓存失败，则记录日志
		l.Logger.Errorf("set user profile to redis failed: %v", err)
	}

	// 6. 返回用户信息
	return &pb.GetUserByAccountResponse{User: &user}, nil
}

// 根据账号查询用户ID
func (l *GetUserByAccountLogic) getUserIdByAccount(account string) (int64, error) {
	var userId int64

	// 1.1 先查询缓存中是否存在用户 id
	// key: user:account:{account}
	// value: user_id
	userId, err := l.svcCtx.Redis.Get(l.ctx, fmt.Sprintf("user:account:%s", account)).Int64()
	// 1.2 如果 Redis 返回缓存命中，则直接返回
	if err == nil && userId != 0 {
		// 如果缓存命中并且 userId 合法，则直接返回
		return userId, nil
	} else if errors.Is(err, redis.Nil) {
		// 1.3 如果 Redis 返回 key 不存在，记日志，继续走数据库
		l.Logger.Infof("get user id from redis not found: %v", err)
	} else {
		// 1.4 如果 Redis 返回其他错误，记日志，继续走数据库
		l.Logger.Errorf("get user id from redis failed: %v", err)
	}

	// 2.1 如果缓存未命中，则查询数据库，仅查询 user_id 字段
	tx := l.svcCtx.DB.Model(&pb.User{}).Where("account = ?", account).Pluck("user_id", &userId)
	if tx.Error != nil {
		// 2.2 如果查询失败，则返回数据库错误
		l.Logger.Errorf("query user in database failed: %v", tx.Error)
		return 0, errors.WithCode(code.CodeDBQueryFailed, "query user in database failed")
	}

	// 2.3 如果没有任何记录，或者 userId 仍为 0，则认为用户不存在
	if tx.RowsAffected == 0 || userId == 0 {
		l.Logger.Errorf("user not found in database by account: %s", account)
		return 0, errors.WithCode(code.CodeUserNotFound, "user not found in database")
	}

	// 3.1 回写缓存
	// key: user:account:{account}
	// value: user_id
	err = l.svcCtx.Redis.Set(l.ctx, fmt.Sprintf("user:account:%s", account), userId, 0).Err()
	if err != nil {
		// 3.2 如果回写缓存失败，则记录日志
		l.Logger.Errorf("set user id to redis failed: %v", err)
	}

	// 4. 返回用户ID
	return userId, nil
}
