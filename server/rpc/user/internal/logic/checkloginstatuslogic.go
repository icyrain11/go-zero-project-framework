package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
	error2 "go-zero-demo/server/common/error"
	common "go-zero-demo/server/common/user"
	"go-zero-demo/server/rpc/user/internal/constant"
	error3 "go-zero-demo/server/rpc/user/internal/error"
	"go-zero-demo/server/rpc/user/internal/svc"
	"go-zero-demo/server/rpc/user/pb/user"
)

type CheckLoginStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLoginStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLoginStatusLogic {
	return &CheckLoginStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLoginStatusLogic) CheckLoginStatus(in *user.CheckLoginStatusReq) (out *user.CheckLoginStatusResp, err error) {

	token := in.Token

	//Token是否有效
	userCtx, err := l.isValidToken(token)
	if err != nil {
		return nil, err
	}

	//刷新token有效期
	err = l.refreshToken(token)
	if err != nil {
		return nil, err
	}

	return &user.CheckLoginStatusResp{
		Id:       userCtx.Id,
		Username: userCtx.Username,
	}, nil
}

func (l *CheckLoginStatusLogic) isValidToken(token string) (userCtx *common.UserCtx, err error) {

	userCtxStr, err := l.svcCtx.Redis.GetCtx(l.ctx, constant.UserLoginKey+token)

	if err != nil {
		l.Logger.Errorf("Redis Get Error %v", err)
		err = errors.New(error2.RedisGetErrorCode, error2.ServerInternalErrorMsg)
		return nil, err
	}

	//未登录
	if userCtxStr == "" {
		err = errors.New(error3.UserNotLoginErrorCode, error3.UserNotLoginErrorMsg)
		return nil, err
	}

	//反序列化
	err = json.Unmarshal([]byte(userCtxStr), userCtx)

	if err != nil {
		l.Logger.Errorf("Json Unmarshal Error %v", err)
		err = errors.New(error2.JsonUnmarshalErrorCode, error2.ServerInternalErrorMsg)
		return nil, err
	}
	return userCtx, nil
}

func (l *CheckLoginStatusLogic) refreshToken(token string) (err error) {
	err = l.svcCtx.Redis.Expire(constant.UserLoginKey+token, constant.UserLoginExpireTime)
	if err != nil {
		l.Logger.Errorf("Redis Expire Error %v", err)
		err = errors.New(error2.RedisExpireErrorCode, error2.ServerInternalErrorMsg)
		return err
	}
	return nil
}
