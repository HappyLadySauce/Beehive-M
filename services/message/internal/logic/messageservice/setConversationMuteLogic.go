package messageservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/message/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/message/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetConversationMuteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetConversationMuteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetConversationMuteLogic {
	return &SetConversationMuteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetConversationMuteLogic) SetConversationMute(in *pb.SetConversationMuteRequest) (*pb.SetConversationMuteResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.SetConversationMuteResponse{}, nil
}
