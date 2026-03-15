package presenceservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/presence/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/presence/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReportHeartbeatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReportHeartbeatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReportHeartbeatLogic {
	return &ReportHeartbeatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReportHeartbeatLogic) ReportHeartbeat(in *pb.ReportHeartbeatRequest) (*pb.ReportHeartbeatResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.ReportHeartbeatResponse{}, nil
}
