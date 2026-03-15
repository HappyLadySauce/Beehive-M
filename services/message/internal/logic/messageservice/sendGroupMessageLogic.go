package messageservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/message/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/message/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendGroupMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendGroupMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendGroupMessageLogic {
	return &SendGroupMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 群聊
func (l *SendGroupMessageLogic) SendGroupMessage(in *pb.SendGroupMessageRequest) (*pb.SendGroupMessageResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.SendGroupMessageResponse{}, nil
}
