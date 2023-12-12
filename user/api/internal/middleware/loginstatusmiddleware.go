package middleware

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/x/errors"
	xhttp "github.com/zeromicro/x/http"
	"net/http"
)

type LoginStatusMiddleware struct {
	Redis *redis.Redis
}

func NewLoginStatusMiddleware(rds *redis.Redis) *LoginStatusMiddleware {
	return &LoginStatusMiddleware{
		Redis: rds,
	}
}

type UserCtx struct {
	Id       int64
	Username string
}

func (m *LoginStatusMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		if token == "" {
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New(400001, "用户未登录"))
			return
		}
		userCtx, err := m.isValidToken(token)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New(400001, "用户未登录"))
			return
		}
		//续签
		err = m.refreshToken(r.Context(), token)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New(400001, "用户未登录"))
			return
		}
		next(w, r.WithContext(context.WithValue(r.Context(), "userCtx", userCtx)))
	}
}

func (m *LoginStatusMiddleware) refreshToken(ctx context.Context, token string) (err error) {
	err = m.Redis.ExpireCtx(ctx, "user:login:"+token, 86400)
	if err != nil {
		return err
	}
	return nil
}

func (m *LoginStatusMiddleware) isValidToken(token string) (userCtx *UserCtx, err error) {
	userCtxStr, err := m.Redis.Get("user:login:" + token)
	if err != nil {
		return nil, err
	}

	if userCtxStr == "" {
		err = errors.New(400001, "用户未登录")
		return nil, err
	}

	err = json.Unmarshal([]byte(userCtxStr), &userCtx)

	if err != nil {
		return nil, err
	}
	return userCtx, nil
}
