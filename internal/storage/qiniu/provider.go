package qiniu

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
	"time"
)

// GetUploadToken 获取uploadToken
func (k *KodoStorage) GetUploadToken(fileAddr string, mimeType string, deadline int64, options *map[string]string) (uploadToken string) {
	scope := fmt.Sprintf("%s:%s", k.Bucket, fileAddr)
	putPolicy := storage.PutPolicy{
		Scope:           scope,
		MimeLimit:       mimeType,
		IsPrefixalScope: 0,
		Expires:         uint64(deadline - time.Now().Unix()),
	}

	if options != nil {
		putPolicy.PersistentOps = (*options)["persistentOps"]
	}

	return putPolicy.UploadToken(k.Mac)
}

// UploadFile 上传文件
func (k *KodoStorage) UploadFile(fileAddr string, input io.Reader, uploadToken string) (err error) {
	config := storage.Config{}
	formUploader := storage.NewFormUploader(&config)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	err = formUploader.Put(context.Background(), &ret, uploadToken, fileAddr, input, -1, &putExtra)
	return err
}
