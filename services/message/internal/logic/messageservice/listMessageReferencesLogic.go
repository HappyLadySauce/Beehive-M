package messageservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/message/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/message/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMessageReferencesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListMessageReferencesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMessageReferencesLogic {
	return &ListMessageReferencesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListMessageReferencesLogic) ListMessageReferences(in *pb.ListMessageReferencesRequest) (*pb.ListMessageReferencesResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.ListMessageReferencesResponse{}, nil
}
