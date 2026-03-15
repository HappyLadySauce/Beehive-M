package messageservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/message/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/message/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrCreateConversationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrCreateConversationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrCreateConversationLogic {
	return &GetOrCreateConversationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrCreateConversationLogic) GetOrCreateConversation(in *pb.GetOrCreateConversationRequest) (*pb.GetOrCreateConversationResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetOrCreateConversationResponse{}, nil
}
