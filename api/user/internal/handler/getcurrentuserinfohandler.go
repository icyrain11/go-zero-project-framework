package handler

import (
	xhttp "github.com/zeromicro/x/http"
	"go-zero-demo/api/user/internal/logic"
	"go-zero-demo/api/user/internal/svc"
	"net/http"
)

func getCurrentUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetCurrentUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetCurrentUserInfo()
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
