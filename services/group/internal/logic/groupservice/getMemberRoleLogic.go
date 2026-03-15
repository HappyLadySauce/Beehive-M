package groupservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/group/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/group/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberRoleLogic {
	return &GetMemberRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMemberRoleLogic) GetMemberRole(in *pb.GetMemberRoleRequest) (*pb.GetMemberRoleResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetMemberRoleResponse{}, nil
}
