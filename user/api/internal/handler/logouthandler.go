package handler

import (
	xhttp "github.com/zeromicro/x/http"
	"go-zero-demo/user/api/internal/logic"
	"go-zero-demo/user/api/internal/svc"
	"net/http"
)

func logoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewLogoutLogic(r.Context(), svcCtx)
		token := r.Header.Get("token")
		err := l.Logout(token)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
		}
	}
}
