package friendservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/friend/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/friend/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFriendsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListFriendsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFriendsLogic {
	return &ListFriendsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 好友
func (l *ListFriendsLogic) ListFriends(in *pb.ListFriendsRequest) (*pb.ListFriendsResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.ListFriendsResponse{}, nil
}
