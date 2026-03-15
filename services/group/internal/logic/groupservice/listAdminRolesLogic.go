package groupservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/group/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/group/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAdminRolesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListAdminRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAdminRolesLogic {
	return &ListAdminRolesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 管理员角色权限
func (l *ListAdminRolesLogic) ListAdminRoles(in *pb.ListAdminRolesRequest) (*pb.ListAdminRolesResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.ListAdminRolesResponse{}, nil
}
