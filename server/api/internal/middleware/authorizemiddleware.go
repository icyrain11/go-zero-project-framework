package middleware

import (
	"net/http"
)

type AuthorizeMiddleware struct {
}

func NewAuthorizeMiddleware() *AuthorizeMiddleware {
	return &AuthorizeMiddleware{}
}

func (m *AuthorizeMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO 鉴权

		next(w, r)
	}
}
