package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-demo/server/api/internal/config"
	"go-zero-demo/server/api/internal/middleware"
	"go-zero-demo/server/rpc/user/userclient"
)

type ServiceContext struct {
	Config      config.Config
	LoginStatus rest.Middleware
	Authorize   rest.Middleware
	User        userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {

	//鉴权中间件
	user := userclient.NewUser(zrpc.MustNewClient(c.User))
	//登录状态检查中间件
	loginStatus := middleware.NewLoginStatusMiddleware(user).Handle
	authorize := middleware.NewAuthorizeMiddleware().Handle

	return &ServiceContext{
		Config:      c,
		Authorize:   authorize,
		User:        user,
		LoginStatus: loginStatus,
	}
}
