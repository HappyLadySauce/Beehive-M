package notifyservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/notify/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/notify/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarkAllReadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMarkAllReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarkAllReadLogic {
	return &MarkAllReadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MarkAllReadLogic) MarkAllRead(in *pb.MarkAllReadRequest) (*pb.MarkAllReadResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.MarkAllReadResponse{}, nil
}
