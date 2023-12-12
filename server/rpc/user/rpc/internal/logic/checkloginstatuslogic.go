package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/x/errors"
	common "go-zero-demo/server/common/user"
	"go-zero-demo/server/rpc/user/rpc/internal/svc"
	"go-zero-demo/server/rpc/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
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
	userCtxStr, err := l.svcCtx.Redis.GetCtx(l.ctx, "user:login:"+token)

	if err != nil {
		err = errors.New(50001, "系统繁忙")
		return nil, err
	}

	userCtx := common.UserCtx{}
	//反序列化
	err = json.Unmarshal([]byte(userCtxStr), &userCtx)
	if err != nil {
		err = errors.New(50001, "系统繁忙")
		return nil, err
	}
	return &user.CheckLoginStatusResp{}, nil
}
