package notifyservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/notify/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/notify/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUnreadCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUnreadCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUnreadCountLogic {
	return &GetUnreadCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUnreadCountLogic) GetUnreadCount(in *pb.GetUnreadCountRequest) (*pb.GetUnreadCountResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUnreadCountResponse{}, nil
}
