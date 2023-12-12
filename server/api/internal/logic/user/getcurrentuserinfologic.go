package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/server/api/internal/svc"
	"go-zero-demo/server/api/internal/types"
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
	return nil, nil
}
