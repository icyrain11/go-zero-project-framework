package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	//TODO 注入rpc依赖
	//用户系统
	User zrpc.RpcClientConf
}
