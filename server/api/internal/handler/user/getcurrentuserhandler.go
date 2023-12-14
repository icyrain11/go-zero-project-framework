package user

import (
	xhttp "github.com/zeromicro/x/http"
	"go-zero-demo/server/api/internal/logic/user"
	"go-zero-demo/server/api/internal/svc"
	"net/http"
)

func GetCurrentUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetCurrentUserLogic(r.Context(), svcCtx)
		token := r.Header.Get("token")
		resp, err := l.GetCurrentUser(token)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
