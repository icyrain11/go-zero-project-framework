package user

import (
	"context"
	"go-zero-demo/server/rpc/user/pb/user"

	"go-zero-demo/server/api/internal/svc"
	"go-zero-demo/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCurrentUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCurrentUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCurrentUserLogic {
	return &GetCurrentUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCurrentUserLogic) GetCurrentUser(token string) (resp *types.GetCurrentUserResp, err error) {
	in := &user.GetCurrentUseReq{
		Token: token,
	}
	out, err := l.svcCtx.User.GetCurrentUser(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return &types.GetCurrentUserResp{
		Id:       out.Id,
		Username: out.Username,
		Nickname: out.Nickname,
		Gender:   out.Gender,
		Mobile:   out.Mobile,
	}, nil
}
