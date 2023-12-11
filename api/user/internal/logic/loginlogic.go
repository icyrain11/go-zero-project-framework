package logic

import (
	"context"
	"github.com/google/uuid"
	"github.com/zeromicro/x/errors"
	error2 "go-zero-demo/api/user/internal/common/error"
	"go-zero-demo/api/user/internal/svc"
	"go-zero-demo/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {

	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)

	//用户不存在
	if err != nil {
		err = errors.New(error2.UserNotFoundCode, error2.UserNotFound)
		return nil, err
	}
	//密码错误
	if user.Password != req.Password {
		err = errors.New(error2.UserPasswordErrorCode, error2.UserPasswordError)
	}

	newUUID, err := uuid.NewUUID()
	token := newUUID.String()
	if err != nil {
		err = errors.New(50001, "服务器繁忙")
		return
	}

	err = l.svcCtx.Redis.SetCtx(l.ctx, "token", token)

	if err != nil {
		err = errors.New(50001, "服务器繁忙")
		return
	}

	//build Token
	return &types.LoginResp{
		Id:       user.Id,
		Username: user.Username,
		Token:    token,
	}, nil
}
