package groupservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/group/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/group/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMyNicknameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMyNicknameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMyNicknameLogic {
	return &UpdateMyNicknameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMyNicknameLogic) UpdateMyNickname(in *pb.UpdateMyNicknameRequest) (*pb.UpdateMyNicknameResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateMyNicknameResponse{}, nil
}
