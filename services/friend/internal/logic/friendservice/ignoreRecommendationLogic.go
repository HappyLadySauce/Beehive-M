package friendservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/friend/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/friend/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IgnoreRecommendationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIgnoreRecommendationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IgnoreRecommendationLogic {
	return &IgnoreRecommendationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IgnoreRecommendationLogic) IgnoreRecommendation(in *pb.IgnoreRecommendationRequest) (*pb.IgnoreRecommendationResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.IgnoreRecommendationResponse{}, nil
}
