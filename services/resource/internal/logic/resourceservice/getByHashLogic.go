package resourceservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/resource/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/resource/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByHashLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetByHashLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByHashLogic {
	return &GetByHashLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetByHashLogic) GetByHash(in *pb.GetByHashRequest) (*pb.GetByHashResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetByHashResponse{}, nil
}
