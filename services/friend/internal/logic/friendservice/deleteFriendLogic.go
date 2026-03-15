package friendservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/friend/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/friend/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFriendLogic {
	return &DeleteFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteFriendLogic) DeleteFriend(in *pb.DeleteFriendRequest) (*pb.DeleteFriendResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.DeleteFriendResponse{}, nil
}
