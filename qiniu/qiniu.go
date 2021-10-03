package qiniu

import (
	"context"
	"fmt"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// 对接文档：https://developer.qiniu.com/kodo/1238/go

type Qiniu struct {
	AccessKey  string
	SecretKey  string
	BucketName string
	Day        int
}

func NewQiniuClient(AccessKey, SecretKey, BucketName string) *Qiniu {
	return &Qiniu{AccessKey, SecretKey, BucketName, 1}
}

// 上传文件
func (q Qiniu) Upload(filename, filepath string) error {

	putPolicy := storage.PutPolicy{
		Scope: q.BucketName,
	}
	mac := qbox.NewMac(q.AccessKey, q.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutFile(context.Background(), &ret, upToken, filename, filepath, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(ret.Key, ret.Hash)
	return nil
}

// 新建bucket管理器
func (q Qiniu) NewBucketManager() *storage.BucketManager {
	mac := qbox.NewMac(q.AccessKey, q.SecretKey)
	cfg := storage.Config{}
	return storage.NewBucketManager(mac, &cfg)
}

// 设置文件生命周期，过一段时间后自动删除文件
func (q Qiniu) DeleteAferDay(key string, day int) error {
	bucketManager := q.NewBucketManager()
	err := bucketManager.DeleteAfterDays(q.BucketName, key, day)
	if err != nil {
		return err
	}
	return nil
}

// 上传文件并设置生命周期
func (q Qiniu) UploadAndDeleteAfter(filename, filepath string) error {
	err := q.Upload(filename, filepath)
	if err != nil {
		return err
	}
	err = q.DeleteAferDay(filename, q.Day)
	if err != nil {
		return err
	}
	return nil
}
