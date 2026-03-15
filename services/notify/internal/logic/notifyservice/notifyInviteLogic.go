package notifyservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/notify/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/notify/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotifyInviteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNotifyInviteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotifyInviteLogic {
	return &NotifyInviteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *NotifyInviteLogic) NotifyInvite(in *pb.NotifyInviteRequest) (*pb.NotifyInviteResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.NotifyInviteResponse{}, nil
}
