package friendservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/friend/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/friend/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListReceivedRequestsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListReceivedRequestsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListReceivedRequestsLogic {
	return &ListReceivedRequestsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListReceivedRequestsLogic) ListReceivedRequests(in *pb.ListReceivedRequestsRequest) (*pb.ListReceivedRequestsResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.ListReceivedRequestsResponse{}, nil
}
