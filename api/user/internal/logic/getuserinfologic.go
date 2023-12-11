package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"go-zero-demo/api/user/internal/svc"
	"go-zero-demo/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)
	if err != nil {
		logc.Errorf(l.ctx, "userModel Operation err %", err)
	}
	fmt.Println(user)
	return
}
