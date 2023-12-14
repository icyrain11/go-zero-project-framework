package logic

import (
	"context"
	"github.com/zeromicro/x/errors"
	"go-zero-demo/server/rpc/item/internal/svc"
	"go-zero-demo/server/rpc/item/pb/item"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUploadTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUploadTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUploadTaskLogic {
	return &CreateUploadTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUploadTaskLogic) CreateUploadTask(in *item.CreateUploadTaskReq) (out *item.CreateUploadTaskResp, err error) {
	t := in.Type
	//TODO 策略模式优化

	if t == "video" {

	} else if t == "cover" {

	} else if t == "avatar" {

	} else {
		err = errors.New(66666, "无实现方法")
	}
	return &item.CreateUploadTaskResp{}, nil
}
