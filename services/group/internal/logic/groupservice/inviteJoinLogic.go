package groupservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/group/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/group/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type InviteJoinLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInviteJoinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InviteJoinLogic {
	return &InviteJoinLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InviteJoinLogic) InviteJoin(in *pb.InviteJoinRequest) (*pb.InviteJoinResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.InviteJoinResponse{}, nil
}
