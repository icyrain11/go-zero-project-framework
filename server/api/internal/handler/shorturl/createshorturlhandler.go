package shorturl

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
	"go-zero-demo/server/api/internal/logic/shorturl"
	"go-zero-demo/server/api/internal/svc"
	"go-zero-demo/server/api/internal/types"
	"net/http"
)

func CreateShortUrlHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateShortUrlReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := shorturl.NewCreateShortUrlLogic(r.Context(), svcCtx)
		err := l.CreateShortUrl(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
		}
	}
}
