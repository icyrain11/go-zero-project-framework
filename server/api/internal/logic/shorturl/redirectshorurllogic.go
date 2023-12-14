package shorturl

import (
	"context"

	"go-zero-demo/server/api/internal/svc"
	"go-zero-demo/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedirectShorUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRedirectShorUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedirectShorUrlLogic {
	return &RedirectShorUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RedirectShorUrlLogic) RedirectShorUrl(req *types.RedirectShorUrlReq) (resp *types.RedirectShorUrlResp, err error) {
	// todo: add your logic here and delete this line
	return
}
