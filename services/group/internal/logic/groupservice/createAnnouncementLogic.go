package groupservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/group/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/group/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAnnouncementLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAnnouncementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAnnouncementLogic {
	return &CreateAnnouncementLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 公告
func (l *CreateAnnouncementLogic) CreateAnnouncement(in *pb.CreateAnnouncementRequest) (*pb.CreateAnnouncementResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.CreateAnnouncementResponse{}, nil
}
