package groupservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/group/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/group/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandleGroupRequestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHandleGroupRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleGroupRequestLogic {
	return &HandleGroupRequestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HandleGroupRequestLogic) HandleGroupRequest(in *pb.HandleGroupRequestRequest) (*pb.HandleGroupRequestResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.HandleGroupRequestResponse{}, nil
}
