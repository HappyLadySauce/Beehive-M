package groupservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/group/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/group/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GrantRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGrantRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrantRoleLogic {
	return &GrantRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GrantRoleLogic) GrantRole(in *pb.GrantRoleRequest) (*pb.GrantRoleResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GrantRoleResponse{}, nil
}
