package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-demo/user/api/internal/config"
	"go-zero-demo/user/api/internal/middleware"
	"go-zero-demo/user/model/mysql/user"
)

type ServiceContext struct {
	Config      config.Config
	UserModel   user.UserModel
	Redis       *redis.Redis
	Authorize   rest.Middleware
	LoginStatus rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	//数据库链接
	sqlConn := sqlx.NewMysql(c.MySql.DataSource)
	rds := redis.MustNewRedis(redis.RedisConf{
		Host: c.Redis.Host,
		Type: c.Redis.Type,
		Pass: c.Redis.Pass,
		Tls:  c.Redis.Tls,
	})

	//鉴权中间件
	loginStatus := middleware.NewLoginStatusMiddleware(rds).Handle

	authorize := middleware.NewAuthorizeMiddleware().Handle

	return &ServiceContext{
		Config:      c,
		UserModel:   user.NewUserModel(sqlConn),
		Authorize:   authorize,
		LoginStatus: loginStatus,
		Redis:       rds,
	}
}
