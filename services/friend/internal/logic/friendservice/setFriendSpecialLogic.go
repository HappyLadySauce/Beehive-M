package friendservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/friend/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/friend/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetFriendSpecialLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetFriendSpecialLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetFriendSpecialLogic {
	return &SetFriendSpecialLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetFriendSpecialLogic) SetFriendSpecial(in *pb.SetFriendSpecialRequest) (*pb.SetFriendSpecialResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.SetFriendSpecialResponse{}, nil
}
