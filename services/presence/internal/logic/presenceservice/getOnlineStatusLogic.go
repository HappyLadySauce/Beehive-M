package presenceservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/presence/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/presence/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOnlineStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOnlineStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOnlineStatusLogic {
	return &GetOnlineStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOnlineStatusLogic) GetOnlineStatus(in *pb.GetOnlineStatusRequest) (*pb.GetOnlineStatusResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetOnlineStatusResponse{}, nil
}
