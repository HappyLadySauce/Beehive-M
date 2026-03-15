package friendservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/friend/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/friend/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckBlacklistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckBlacklistLogic {
	return &CheckBlacklistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckBlacklistLogic) CheckBlacklist(in *pb.CheckBlacklistRequest) (*pb.CheckBlacklistResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.CheckBlacklistResponse{}, nil
}
