package middleware

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/x/errors"
	xhttp "github.com/zeromicro/x/http"
	"net/http"
)

type AuthorizeHandlerMiddleware struct {
	Redis *redis.Redis
}

func NewAuthorizeHandlerMiddleware(rds *redis.Redis) *AuthorizeHandlerMiddleware {
	return &AuthorizeHandlerMiddleware{
		Redis: rds,
	}
}

func (m *AuthorizeHandlerMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//token
		token := r.Header.Get("token")
		userCtx, err := m.isValidToken(token)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New(400001, "用户未登录"))
		} else {
			next(w, r.WithContext(context.WithValue(r.Context(), "userCtx", userCtx)))
		}

	}
}

type UserCtx struct {
	Id       int64
	Username string
}

func (m *AuthorizeHandlerMiddleware) isValidToken(token string) (userCtx *UserCtx, err error) {
	userCtxStr, err := m.Redis.Get("user:login:" + token)
	if err != nil {
		return nil, err
	}

	if userCtxStr == "" {
		return nil, err
	}

	err = json.Unmarshal([]byte(userCtxStr), &userCtx)

	if err != nil {
		return nil, err
	}

	return userCtx, nil
}
