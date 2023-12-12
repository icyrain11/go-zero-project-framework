package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	//TODO 注入rpc依赖
}
