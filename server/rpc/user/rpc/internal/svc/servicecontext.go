package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/server/rpc/user/model/mysql/user"
	"go-zero-demo/server/rpc/user/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	Redis     *redis.Redis
	UserModel user.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	//链接数据库
	sqlConn := sqlx.NewMysql(c.MySql.DataSource)
	//链接Redis
	rds := redis.MustNewRedis(c.Redis.RedisConf)
	return &ServiceContext{
		Config:    c,
		Redis:     rds,
		UserModel: user.NewUserModel(sqlConn),
	}
}
