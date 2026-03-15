package userservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/user/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserBatchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserBatchLogic {
	return &GetUserBatchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量按 user_id 查（会话列表、群成员列表等）
func (l *GetUserBatchLogic) GetUserBatch(in *pb.GetUserBatchRequest) (*pb.GetUserBatchResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserBatchResponse{}, nil
}
