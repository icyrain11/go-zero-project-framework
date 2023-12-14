package logic

import (
	"context"
	"github.com/zeromicro/x/errors"
	error2 "go-zero-demo/server/rpc/item/internal/error"

	"go-zero-demo/server/rpc/item/internal/svc"
	"go-zero-demo/server/rpc/item/pb/item"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoByIdLogic {
	return &GetVideoByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoByIdLogic) GetVideoById(in *item.GetVideoByIdReq) (*item.GetVideoByIdResp, error) {
	id := in.Id

	//查询视频信息
	one, err := l.svcCtx.VideoModel.FindOne(l.ctx, id)
	if err != nil {
		l.Logger.Errorf("Get Video Error %v", err)
		err = errors.New(error2.VideoNotFoundErrorCode, error2.VideoNotFoundErrorMsg)
		return nil, err
	}

	return &item.GetVideoByIdResp{
		Id:          one.Id,
		UserId:      one.UserId,
		Status:      one.Status,
		Description: one.Description,
		FileAddr:    one.FileAddr,
		CoverAddr:   one.CoverAddr,
		JobId:       one.JobId,
	}, nil
}
