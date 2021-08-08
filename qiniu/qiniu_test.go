package qiniu

import "testing"

func TestUploadToQiniu(t *testing.T) {
	filepath := "/Users/tao/Documents/workspace/go/gomod/MdHelper/README.assets/image-20210808202426899.png"
	filename := "github-x.png"

	qiniuClient := Qiniu{
		AccessKey:  "your AccessKey",
		SecretKey:  "your SecretKey",
		BucketName: "your Bucket",
	}

	qiniuClient.UploadAndDeleteAfter(filename, filepath, 1)

}
