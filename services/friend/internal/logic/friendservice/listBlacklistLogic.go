package friendservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/friend/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/friend/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListBlacklistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListBlacklistLogic {
	return &ListBlacklistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListBlacklistLogic) ListBlacklist(in *pb.ListBlacklistRequest) (*pb.ListBlacklistResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.ListBlacklistResponse{}, nil
}
