package notifyservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/notify/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/notify/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotifyGroupRequestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNotifyGroupRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotifyGroupRequestLogic {
	return &NotifyGroupRequestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *NotifyGroupRequestLogic) NotifyGroupRequest(in *pb.NotifyGroupRequestRequest) (*pb.NotifyGroupRequestResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.NotifyGroupRequestResponse{}, nil
}
