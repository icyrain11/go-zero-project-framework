package item

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
	"go-zero-demo/server/api/internal/logic/item"
	"go-zero-demo/server/api/internal/svc"
	"go-zero-demo/server/api/internal/types"
	"net/http"
)

func CreateUploadTaskHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateUplaodTaskReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := item.NewCreateUploadTaskLogic(r.Context(), svcCtx)
		resp, err := l.CreateUploadTask(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
