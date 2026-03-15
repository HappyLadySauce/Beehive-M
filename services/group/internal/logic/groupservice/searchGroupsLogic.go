package groupservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/group/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/group/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchGroupsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchGroupsLogic {
	return &SearchGroupsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchGroupsLogic) SearchGroups(in *pb.SearchGroupsRequest) (*pb.SearchGroupsResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchGroupsResponse{}, nil
}
