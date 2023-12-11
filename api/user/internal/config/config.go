package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	MySql struct {
		DataSource string
	}
	Auth struct { // JWT 认证需要的密钥和过期时间配置
		AccessSecret string
		AccessExpire int64
	}
	Redis struct {
		Host string
		Type string `json:",default=node,options=node|cluster"`
		Pass string `json:",optional"`
		Tls  bool   `json:",optional"`
	}
}
