package userservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/user/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
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
func (l *GetUserLogic) GetUser(in *pb.GetUserRequest) (*pb.User, error) {
	// todo: add your logic here and delete this line

	return &pb.User{}, nil
}
