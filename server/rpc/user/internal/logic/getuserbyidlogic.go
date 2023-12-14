package logic

import (
	"context"
	"github.com/zeromicro/x/errors"
	error2 "go-zero-demo/server/rpc/user/internal/error"

	"go-zero-demo/server/rpc/user/internal/svc"
	"go-zero-demo/server/rpc/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *user.GetUserByIdReq) (*user.GetUserByIdResp, error) {
	id := in.Id

	one, err := l.svcCtx.UserModel.FindOne(l.ctx, id)
	if err != nil {
		err = errors.New(error2.UserNotFoundErrorCode, error2.UserNotFoundErrorMsg)
		return nil, err
	}

	return &user.GetUserByIdResp{
		Id:       one.Id,
		Username: one.Username,
		Nickname: one.Nickname,
		Gender:   one.Gender,
		Mobile:   one.Mobile,
	}, nil
}
