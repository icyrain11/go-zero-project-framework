package item

import (
	"context"

	"go-zero-demo/server/api/internal/svc"
	"go-zero-demo/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUploadTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUploadTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUploadTaskLogic {
	return &CreateUploadTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUploadTaskLogic) CreateUploadTask(req *types.CreateUplaodTaskReq) (resp *types.CreateUploadTaskResp, err error) {
	// todo: add your logic here and delete this line

	return
}
