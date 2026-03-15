package groupservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/group/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/group/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarkAnnouncementReadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMarkAnnouncementReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarkAnnouncementReadLogic {
	return &MarkAnnouncementReadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MarkAnnouncementReadLogic) MarkAnnouncementRead(in *pb.MarkAnnouncementReadRequest) (*pb.MarkAnnouncementReadResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.MarkAnnouncementReadResponse{}, nil
}
