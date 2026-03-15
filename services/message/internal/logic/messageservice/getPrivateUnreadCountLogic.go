package messageservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/message/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/message/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrivateUnreadCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPrivateUnreadCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrivateUnreadCountLogic {
	return &GetPrivateUnreadCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPrivateUnreadCountLogic) GetPrivateUnreadCount(in *pb.GetPrivateUnreadCountRequest) (*pb.GetPrivateUnreadCountResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetPrivateUnreadCountResponse{}, nil
}
