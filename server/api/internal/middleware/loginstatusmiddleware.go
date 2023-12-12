package middleware

import (
	"context"
	xhttp "github.com/zeromicro/x/http"
	common "go-zero-demo/server/common/user"
	"go-zero-demo/server/rpc/user/pb/user"
	"go-zero-demo/server/rpc/user/userclient"
	"net/http"
)

type LoginStatusMiddleware struct {
	User userclient.User
}

func NewLoginStatusMiddleware(user userclient.User) *LoginStatusMiddleware {
	return &LoginStatusMiddleware{
		User: user,
	}
}

func (m *LoginStatusMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		in := &user.CheckLoginStatusReq{
			Token: token,
		}
		out, err := m.User.CheckLoginStatus(r.Context(), in)
		if err != nil {
			xhttp.JsonBaseResponse(w, err)
			return
		}

		userCtx := common.UserCtx{
			Id:       out.UserCtx.Id,
			Username: out.UserCtx.Username,
		}

		next(w, r.WithContext(context.WithValue(r.Context(), "userCtx", userCtx)))

	}
}
