package presenceservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/presence/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/presence/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchGetOnlineStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchGetOnlineStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchGetOnlineStatusLogic {
	return &BatchGetOnlineStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BatchGetOnlineStatusLogic) BatchGetOnlineStatus(in *pb.BatchGetOnlineStatusRequest) (*pb.BatchGetOnlineStatusResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.BatchGetOnlineStatusResponse{}, nil
}
