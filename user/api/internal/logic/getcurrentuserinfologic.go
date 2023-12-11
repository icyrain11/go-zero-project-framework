package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
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
	fmt.Println(userCtx.Id)
	return
}
