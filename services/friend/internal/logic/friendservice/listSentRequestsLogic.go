package friendservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/friend/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/friend/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSentRequestsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListSentRequestsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSentRequestsLogic {
	return &ListSentRequestsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListSentRequestsLogic) ListSentRequests(in *pb.ListSentRequestsRequest) (*pb.ListSentRequestsResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.ListSentRequestsResponse{}, nil
}
