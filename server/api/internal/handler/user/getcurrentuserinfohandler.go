package user

import (
	xhttp "github.com/zeromicro/x/http"
	"go-zero-demo/server/api/internal/logic/user"
	"go-zero-demo/server/api/internal/svc"
	"net/http"
)

func GetCurrentUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetCurrentUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetCurrentUserInfo()
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
