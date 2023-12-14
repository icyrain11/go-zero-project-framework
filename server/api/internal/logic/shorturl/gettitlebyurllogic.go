package shorturl

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/zeromicro/go-zero/rest/httpc"
	"github.com/zeromicro/x/errors"
	error3 "go-zero-demo/internal/error"
	"go-zero-demo/server/api/internal/svc"
	"go-zero-demo/server/api/internal/types"
	"io"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTitleByUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTitleByUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTitleByUrlLogic {
	return &GetTitleByUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTitleByUrlLogic) GetTitleByUrl(req *types.GetTitleByUrlReq) (resp *types.GetTitleByUrlResp, err error) {
	url := req.Url
	do, err := httpc.Do(l.ctx, http.MethodGet, url, nil)
	if err != nil {
		err = errors.New(error3.HttpGetRequestError, error3.ServerInternalErrorMsg)
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			l.Logger.Errorf("Body Close Error %v", err)
		}
	}(do.Body)

	//成功
	if do.StatusCode == http.StatusOK {
		doc, err := goquery.NewDocumentFromReader(do.Body)
		if err != nil {
			err = errors.New(error3.DocumentFromReaderErrorCode, error3.ServerInternalErrorMsg)
		}
		//返回title文本
		title := doc.Find("title").Text()
		return &types.GetTitleByUrlResp{
			Title: title,
		}, nil
	}

	return &types.GetTitleByUrlResp{
		Title: "Error while fetching title.",
	}, nil
}
