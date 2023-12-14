package logic

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
	user3 "go-zero-demo/internal/model/mysql/user"
	"go-zero-demo/internal/user"
	error2 "go-zero-demo/server/rpc/user/internal/error"
	"go-zero-demo/server/rpc/user/internal/svc"
	"go-zero-demo/server/rpc/user/pb/user"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (out *user.LoginResp, err error) {

	one, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)

	//用户不存在
	if err != nil {
		err = errors.New(error2.UserNotFoundErrorCode, error2.UserNotFoundErrorMsg)
		return nil, err
	}

	//密码错误
	if one.Password != in.Password {
		err = errors.New(error2.UserPasswordErrorCode, error2.UserPasswordErrorMsg)
		return nil, err
	}

	token, err := l.buildAndSetToken(one)

	if err != nil {
		return nil, err
	}

	return &user.LoginResp{
		Id:       one.Id,
		Username: one.Username,
		Token:    token,
	}, nil
}

func (l *LoginLogic) buildAndSetToken(user *user3.User) (token string, err error) {
	//build uuid
	newUUID, err := uuid.NewUUID()

	token = newUUID.String()
	if err != nil {
		l.Logger.Errorf("UUID Generate Error %v", err)
		err = errors.New(50001, "服务器繁忙")
		return "", err
	}

	userCtx := common.UserCtx{
		Username: user.Username,
		Id:       user.Id,
	}

	userCtxStr, err := json.Marshal(userCtx)

	err = l.svcCtx.Redis.SetexCtx(l.ctx, "user:login:"+token, string(userCtxStr), 86400)

	if err != nil {
		l.Logger.Errorf("Redis SetEx Error %v", err)
		return "", err
	}

	return token, nil
}
