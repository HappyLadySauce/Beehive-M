package messageservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/message/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/message/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarkPrivateReadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMarkPrivateReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarkPrivateReadLogic {
	return &MarkPrivateReadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MarkPrivateReadLogic) MarkPrivateRead(in *pb.MarkPrivateReadRequest) (*pb.MarkPrivateReadResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.MarkPrivateReadResponse{}, nil
}
