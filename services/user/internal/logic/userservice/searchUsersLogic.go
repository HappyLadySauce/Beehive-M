package userservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/user/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUsersLogic {
	return &SearchUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 关键词搜索用户（全文检索），分页
func (l *SearchUsersLogic) SearchUsers(in *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchUsersResponse{}, nil
}
