package resourceservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/resource/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/resource/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CleanExpiredTempLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCleanExpiredTempLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CleanExpiredTempLogic {
	return &CleanExpiredTempLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CleanExpiredTempLogic) CleanExpiredTemp(in *pb.CleanExpiredTempRequest) (*pb.CleanExpiredTempResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.CleanExpiredTempResponse{}, nil
}
