package groupservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/group/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/group/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetMemberMutedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetMemberMutedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetMemberMutedLogic {
	return &SetMemberMutedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetMemberMutedLogic) SetMemberMuted(in *pb.SetMemberMutedRequest) (*pb.SetMemberMutedResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.SetMemberMutedResponse{}, nil
}
