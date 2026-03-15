package groupservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/group/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/group/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListGroupRequestsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListGroupRequestsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListGroupRequestsLogic {
	return &ListGroupRequestsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListGroupRequestsLogic) ListGroupRequests(in *pb.ListGroupRequestsRequest) (*pb.ListGroupRequestsResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.ListGroupRequestsResponse{}, nil
}
