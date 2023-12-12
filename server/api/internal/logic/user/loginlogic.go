package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/server/api/internal/svc"
	"go-zero-demo/server/api/internal/types"
	"go-zero-demo/server/rpc/user/pb/user"
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
	//调用rpc
	in := &user.LoginReq{
		Username: req.Username,
		Password: req.Password,
	}

	out, err := l.svcCtx.User.Login(l.ctx, in)

	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		Id:       out.Id,
		Username: out.Username,
		Token:    out.Token,
	}, nil
}
