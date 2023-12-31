package logic

import (
	"context"
	"fmt"
	"go-zero-demo/internal/user"
	"go-zero-demo/server/rpc/user/internal/svc"
	"go-zero-demo/server/rpc/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCurrentUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCurrentUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCurrentUserLogic {
	return &GetCurrentUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCurrentUserLogic) GetCurrentUser(in *user.GetCurrentUseReq) (*user.GetCurrentUserResp, error) {
	// todo: add your logic here and delete this line
	userCtx := l.ctx.Value("userCtx").(*common.UserCtx)
	fmt.Println(userCtx)
	return &user.GetCurrentUserResp{}, nil
}
