package friendservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/friend/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/friend/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListRecommendationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListRecommendationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListRecommendationsLogic {
	return &ListRecommendationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 推荐
func (l *ListRecommendationsLogic) ListRecommendations(in *pb.ListRecommendationsRequest) (*pb.ListRecommendationsResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.ListRecommendationsResponse{}, nil
}
