package handler

import (
	"net/http"
	"github.com/zeromicro/go-zero/rest/httpx"
	"xhttp "github.com/zeromicro/x/http""
	"go-zero-demo/relation/internal/logic"
	"go-zero-demo/relation/internal/svc"
	"go-zero-demo/relation/internal/types"


)

func RelationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRelationLogic(r.Context(), svcCtx)
		resp, err := l.Relation(&req)
        if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
