package userservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/user/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserForAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserForAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserForAuthLogic {
	return &GetUserForAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 仅内部使用：按 user_id 查用户（含 password_hash），供 auth 登录校验
func (l *GetUserForAuthLogic) GetUserForAuth(in *pb.GetUserForAuthRequest) (*pb.UserForAuth, error) {
	// todo: add your logic here and delete this line

	return &pb.UserForAuth{}, nil
}
