package messageservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/message/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/message/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSyncCursorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSyncCursorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSyncCursorLogic {
	return &GetSyncCursorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 同步
func (l *GetSyncCursorLogic) GetSyncCursor(in *pb.GetSyncCursorRequest) (*pb.GetSyncCursorResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetSyncCursorResponse{}, nil
}
