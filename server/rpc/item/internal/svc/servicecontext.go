package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/internal/model/mysql/uploadtask"
	"go-zero-demo/internal/model/mysql/video"
	"go-zero-demo/server/rpc/item/internal/config"
)

type ServiceContext struct {
	Config          config.Config
	Redis           *redis.Redis
	UploadTaskModel uploadtask.UploadTaskModel
	VideoModel      video.VideoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	//链接数据库
	sqlConn := sqlx.NewMysql(c.MySql.DataSource)
	//链接Redis
	rds := redis.MustNewRedis(c.Redis.RedisConf)
	//TODO 新建一个七牛云OSS
	return &ServiceContext{
		Config:          c,
		Redis:           rds,
		UploadTaskModel: uploadtask.NewUploadTaskModel(sqlConn),
		VideoModel:      video.NewVideoModel(sqlConn),
	}
}
