package handler

import (
	"go-zero-demo/api/common/response"
	"go-zero-demo/api/user/internal/logic"
	"go-zero-demo/api/user/internal/svc"
	"net/http"
)

func getCurrentUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetCurrentUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetCurrentUserInfo()
		response.Response(w, resp, err)
	}
}
