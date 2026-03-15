package notifyservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/notify/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/notify/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotifyFriendRequestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNotifyFriendRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotifyFriendRequestLogic {
	return &NotifyFriendRequestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 内部创建（供 Friend/Group 等调）
func (l *NotifyFriendRequestLogic) NotifyFriendRequest(in *pb.NotifyFriendRequestRequest) (*pb.NotifyFriendRequestResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.NotifyFriendRequestResponse{}, nil
}
