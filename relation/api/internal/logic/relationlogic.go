package logic

import (
	"context"
	"go-zero-demo/relation/api/internal/svc"
	"go-zero-demo/relation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationLogic {
	return &RelationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationLogic) Relation(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
