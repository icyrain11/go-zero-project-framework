package qiniu

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
)

type (
	// KodoStorage 七牛云存储客户端
	KodoStorage struct {
		Mac    *qbox.Mac
		Bucket string
	}
	// KodoStorageConf 七牛云存储配置
	KodoStorageConf struct {
		AccessKey string
		SecretKey string
		Bucket    string
	}
)

// NewKodoStorage 创建七牛云存储实例
func NewKodoStorage(conf KodoStorageConf) (k *KodoStorage) {
	//创建七牛云凭证
	mac := qbox.NewMac(conf.AccessKey, conf.SecretKey)
	return &KodoStorage{
		Mac:    mac,
		Bucket: conf.Bucket,
	}
}
