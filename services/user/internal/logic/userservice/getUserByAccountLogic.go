package userservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/user/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return &pb.GetUserByAccountResponse{}, nil
}
