package friendservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/friend/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/friend/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetFriendMutedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetFriendMutedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetFriendMutedLogic {
	return &SetFriendMutedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetFriendMutedLogic) SetFriendMuted(in *pb.SetFriendMutedRequest) (*pb.SetFriendMutedResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.SetFriendMutedResponse{}, nil
}
