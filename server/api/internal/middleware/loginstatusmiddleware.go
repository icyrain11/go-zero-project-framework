package middleware

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/x/errors"
	xhttp "github.com/zeromicro/x/http"
	common "go-zero-demo/server/common/user"
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

func (m *LoginStatusMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		if token == "" {
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New(400001, "用户未登录"))
			return
		}

		//续签
		err := m.refreshToken(r.Context(), token)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New(400001, "用户未登录"))
			return
		}

		userCtx := common.UserCtx{}
		//获取用户内容

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
