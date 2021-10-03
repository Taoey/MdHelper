package qiniu

import "testing"

func TestUploadToQiniu(t *testing.T) {
	filepath := "/Users/tao/Documents/workspace/go/gomod/MdHelper/README.assets/image-20210808202426899.png"
	filename := "github-x.png"

	qiniuClient := Qiniu{
		AccessKey:  "OcZiLE2w5HFn-piCno178aJ0g32pfI18bG4rjC7v",
		SecretKey:  "CrJjANFUyfSEFp36UsFXsi0DBiuZWL4H6rh_RIM5",
		BucketName: "python_wx_pic",
		Day:        1,
	}

	qiniuClient.UploadAndDeleteAfter(filename, filepath)

}
