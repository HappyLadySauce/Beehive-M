package notifyservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/notify/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/notify/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotifySystemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNotifySystemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotifySystemLogic {
	return &NotifySystemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *NotifySystemLogic) NotifySystem(in *pb.NotifySystemRequest) (*pb.NotifySystemResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.NotifySystemResponse{}, nil
}
