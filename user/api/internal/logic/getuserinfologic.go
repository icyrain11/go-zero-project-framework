package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
	error2 "go-zero-demo/user/api/internal/error"
	"go-zero-demo/user/api/internal/svc"
	"go-zero-demo/user/api/internal/types"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)
	//用户不存在
	if err != nil {
		err = errors.New(error2.UserNotFoundErrorCode, error2.UserNotFoundErrorMsg)
		l.Logger.Errorf("UserModel Operation Error %v", err)
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
