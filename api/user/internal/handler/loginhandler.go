package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/api/common/response"
	"go-zero-demo/api/user/internal/logic"
	"go-zero-demo/api/user/internal/svc"
	"go-zero-demo/api/user/internal/types"
	"net/http"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Response(w, resp, err)
	}
}
