package groupservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/group/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/group/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandleInvitationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHandleInvitationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleInvitationLogic {
	return &HandleInvitationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HandleInvitationLogic) HandleInvitation(in *pb.HandleInvitationRequest) (*pb.HandleInvitationResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.HandleInvitationResponse{}, nil
}
