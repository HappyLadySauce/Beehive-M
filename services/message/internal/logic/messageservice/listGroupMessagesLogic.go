package messageservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/message/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/message/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListGroupMessagesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListGroupMessagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListGroupMessagesLogic {
	return &ListGroupMessagesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListGroupMessagesLogic) ListGroupMessages(in *pb.ListGroupMessagesRequest) (*pb.ListGroupMessagesResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.ListGroupMessagesResponse{}, nil
}
