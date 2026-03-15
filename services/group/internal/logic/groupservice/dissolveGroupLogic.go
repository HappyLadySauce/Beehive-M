package groupservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/group/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/group/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DissolveGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDissolveGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DissolveGroupLogic {
	return &DissolveGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DissolveGroupLogic) DissolveGroup(in *pb.DissolveGroupRequest) (*pb.DissolveGroupResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.DissolveGroupResponse{}, nil
}
