package resourceservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/resource/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/resource/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteResourceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteResourceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteResourceLogic {
	return &DeleteResourceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteResourceLogic) DeleteResource(in *pb.DeleteResourceRequest) (*pb.DeleteResourceResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.DeleteResourceResponse{}, nil
}
