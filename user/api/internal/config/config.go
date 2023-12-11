package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	MySql struct {
		DataSource string
	}
	Redis struct {
		Host string
		Type string `json:",default=node,options=node|cluster"`
		Pass string `json:",optional"`
		Tls  bool   `json:",optional"`
	}
}
