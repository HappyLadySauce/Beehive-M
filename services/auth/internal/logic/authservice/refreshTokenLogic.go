package authservicelogic

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/auth/internal/svc"
	"github.com/HappyLadySauce/Beehive-M/services/auth/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefreshTokenLogic) RefreshToken(in *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.RefreshTokenResponse{}, nil
}
