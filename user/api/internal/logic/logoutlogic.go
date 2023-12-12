package logic

import (
	"context"
	"github.com/zeromicro/x/errors"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/user/api/internal/svc"
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

	del, err := l.svcCtx.Redis.Del("user:login:" + token)

	if err != nil {
		err = errors.New(40001, "用户未登录")
		return err
	}

	if del < 1 {
		err = errors.New(40001, "用户未登录")
		return err
	}

	return nil
}
