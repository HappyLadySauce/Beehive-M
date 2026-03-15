package groupservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/group/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/group/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyJoinLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewApplyJoinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyJoinLogic {
	return &ApplyJoinLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 申请/邀请
func (l *ApplyJoinLogic) ApplyJoin(in *pb.ApplyJoinRequest) (*pb.ApplyJoinResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.ApplyJoinResponse{}, nil
}
