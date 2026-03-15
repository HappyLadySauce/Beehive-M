package friendservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/friend/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/friend/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendFriendRequestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendFriendRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendFriendRequestLogic {
	return &SendFriendRequestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 申请
func (l *SendFriendRequestLogic) SendFriendRequest(in *pb.SendFriendRequestRequest) (*pb.SendFriendRequestResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.SendFriendRequestResponse{}, nil
}
