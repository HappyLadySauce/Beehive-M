package messageservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/message/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/message/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PullMessagesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPullMessagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PullMessagesLogic {
	return &PullMessagesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PullMessagesLogic) PullMessages(in *pb.PullMessagesRequest) (*pb.PullMessagesResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.PullMessagesResponse{}, nil
}
