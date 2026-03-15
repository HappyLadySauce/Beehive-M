package resourceservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/resource/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/resource/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetResourceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetResourceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetResourceLogic {
	return &GetResourceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetResourceLogic) GetResource(in *pb.GetResourceRequest) (*pb.GetResourceResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetResourceResponse{}, nil
}
