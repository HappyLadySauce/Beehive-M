package groupservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/group/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/group/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMyGroupsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListMyGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMyGroupsLogic {
	return &ListMyGroupsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListMyGroupsLogic) ListMyGroups(in *pb.ListMyGroupsRequest) (*pb.ListMyGroupsResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.ListMyGroupsResponse{}, nil
}
