package groupservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/group/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/group/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupLogic {
	return &UpdateGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGroupLogic) UpdateGroup(in *pb.UpdateGroupRequest) (*pb.UpdateGroupResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateGroupResponse{}, nil
}
