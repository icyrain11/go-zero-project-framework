package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
	"go-zero-demo/user/api/internal/middleware"
	"go-zero-demo/user/api/internal/svc"
	"go-zero-demo/user/api/internal/types"
)

type GetCurrentUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCurrentUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCurrentUserInfoLogic {
	return &GetCurrentUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCurrentUserInfoLogic) GetCurrentUserInfo() (resp *types.GetUserInfoResp, err error) {
	userCtx := l.ctx.Value("userCtx").(*middleware.UserCtx)
	id := userCtx.Id
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, id)

	if err != nil {
		err = errors.New(50001, "系统繁忙")
		return nil, err
	}

	return &types.GetUserInfoResp{
		Id:       user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
		Gender:   user.Gender,
		Mobile:   user.Mobile,
	}, nil
}
