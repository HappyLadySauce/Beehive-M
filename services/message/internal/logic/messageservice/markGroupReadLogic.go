package messageservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/message/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/message/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarkGroupReadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMarkGroupReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarkGroupReadLogic {
	return &MarkGroupReadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MarkGroupReadLogic) MarkGroupRead(in *pb.MarkGroupReadRequest) (*pb.MarkGroupReadResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.MarkGroupReadResponse{}, nil
}
