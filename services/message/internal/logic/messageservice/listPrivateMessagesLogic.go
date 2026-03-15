package messageservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/message/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/message/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPrivateMessagesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPrivateMessagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPrivateMessagesLogic {
	return &ListPrivateMessagesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListPrivateMessagesLogic) ListPrivateMessages(in *pb.ListPrivateMessagesRequest) (*pb.ListPrivateMessagesResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.ListPrivateMessagesResponse{}, nil
}
