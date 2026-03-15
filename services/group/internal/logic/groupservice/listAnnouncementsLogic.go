package groupservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/group/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/group/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAnnouncementsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListAnnouncementsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAnnouncementsLogic {
	return &ListAnnouncementsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListAnnouncementsLogic) ListAnnouncements(in *pb.ListAnnouncementsRequest) (*pb.ListAnnouncementsResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.ListAnnouncementsResponse{}, nil
}
