package groupservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/group/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/group/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListInvitationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListInvitationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListInvitationsLogic {
	return &ListInvitationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListInvitationsLogic) ListInvitations(in *pb.ListInvitationsRequest) (*pb.ListInvitationsResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.ListInvitationsResponse{}, nil
}
