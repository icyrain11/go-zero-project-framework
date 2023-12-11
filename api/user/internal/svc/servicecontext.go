package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/api/user/internal/config"
	"go-zero-demo/model/mysql/user"
)

type ServiceContext struct {
	Config    config.Config
	UserModel user.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.MySql.DataSource)

	return &ServiceContext{
		Config:    c,
		UserModel: user.NewUserModel(sqlConn),
	}
}
