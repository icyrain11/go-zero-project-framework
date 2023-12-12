package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"go-zero-demo/server/api/internal/config"
	"go-zero-demo/server/api/internal/middleware"
	user2 "go-zero-demo/server/rpc/user/rpc/pb/user"
)

type ServiceContext struct {
	Config      config.Config
	Authorize   rest.Middleware
	LoginStatus rest.Middleware
	user2.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {

	//鉴权中间件
	authorize := middleware.NewAuthorizeMiddleware().Handle

	return &ServiceContext{
		Config:    c,
		Authorize: authorize,
	}
}
