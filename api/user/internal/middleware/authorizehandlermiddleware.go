package middleware

import (
	"github.com/zeromicro/x/errors"
	xhttp "github.com/zeromicro/x/http"
	"net/http"
)

type AuthorizeHandlerMiddleware struct {
}

func NewAuthorizeHandlerMiddleware() *AuthorizeHandlerMiddleware {
	return &AuthorizeHandlerMiddleware{}
}

func (m *AuthorizeHandlerMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//token
		token := r.Header.Get("token")
		if err := isValidToken(token); err != nil {
			next(w, r)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New(400001, "用户未登录"))
		}
	}
}

func isValidToken(token string) (err error) {
	//TODO 先这么做吧
	if err == nil {
		return err
	}
	//TODO 需要修改
	if token == "" {
		return err
	}
	return nil
}
