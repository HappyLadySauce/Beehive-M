package messageservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/message/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/message/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupUnreadCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupUnreadCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupUnreadCountLogic {
	return &GetGroupUnreadCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupUnreadCountLogic) GetGroupUnreadCount(in *pb.GetGroupUnreadCountRequest) (*pb.GetGroupUnreadCountResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetGroupUnreadCountResponse{}, nil
}
