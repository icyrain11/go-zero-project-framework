package logic

import (
	"context"

	"go-zero-demo/server/rpc/user/internal/svc"
	"go-zero-demo/server/rpc/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogoutLogic) Logout(in *user.LoginReq) (*user.LogoutResp, error) {
	// todo: add your logic here and delete this line

	return &user.LogoutResp{}, nil
}
