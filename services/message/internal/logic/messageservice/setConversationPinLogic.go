package messageservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/message/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/message/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetConversationPinLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetConversationPinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetConversationPinLogic {
	return &SetConversationPinLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetConversationPinLogic) SetConversationPin(in *pb.SetConversationPinRequest) (*pb.SetConversationPinResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.SetConversationPinResponse{}, nil
}
