package user

import (
	"context"
	"go-zero-demo/server/rpc/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/server/api/internal/svc"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(token string) error {
	in := &user.LogoutReq{
		Token: token,
	}
	//退出登录
	_, err := l.svcCtx.User.Logout(l.ctx, in)
	if err != nil {
		return err
	}
	return nil
}
