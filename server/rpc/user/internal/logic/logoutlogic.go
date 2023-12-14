package logic

import (
	"context"
	"github.com/zeromicro/x/errors"
	error2 "go-zero-demo/internal/error"
	"go-zero-demo/server/rpc/user/internal/constant"
	error3 "go-zero-demo/server/rpc/user/internal/error"
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

func (l *LogoutLogic) Logout(in *user.LogoutReq) (*user.LogoutResp, error) {
	token := in.Token

	del, err := l.svcCtx.Redis.Del(constant.UserLoginKey + token)
	if err != nil {
		err = errors.New(error2.RedisDelErrorCode, error2.ServerInternalErrorMsg)
		l.Logger.Errorf("Redis Del Error %v", err)
		return nil, err
	}

	if del < 1 {
		err = errors.New(error3.UserNotLoginErrorCode, error3.UserNotLoginErrorMsg)
		return nil, err
	}

	return &user.LogoutResp{}, nil
}
