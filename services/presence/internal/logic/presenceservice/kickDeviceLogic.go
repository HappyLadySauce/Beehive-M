package presenceservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/presence/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/presence/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type KickDeviceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKickDeviceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KickDeviceLogic {
	return &KickDeviceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *KickDeviceLogic) KickDevice(in *pb.KickDeviceRequest) (*pb.KickDeviceResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.KickDeviceResponse{}, nil
}
