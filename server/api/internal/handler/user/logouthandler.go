package user

import (
	"net/http"
	"github.com/zeromicro/go-zero/rest/httpx"
	"xhttp "github.com/zeromicro/x/http""
	"go-zero-demo/server/api/internal/logic/user"
	"go-zero-demo/server/api/internal/svc"

)

func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewLogoutLogic(r.Context(), svcCtx)
		err := l.Logout()
        if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
