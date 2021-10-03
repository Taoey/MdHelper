package service

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/Taoey/MdHelper/qiniu"
)

type Picture struct {
	Name        string // 文件名
	AbFilePath  string //绝对路径
	QiniuClient *qiniu.Qiniu
}

func NewPicture(uploder *qiniu.Qiniu, name, filepath string) *Picture {
	return &Picture{name, filepath, uploder}
}

func (p *Picture) Upload() {
	p.QiniuClient.UploadAndDeleteAfter(p.Name, p.AbFilePath)
}

var GQiniuClient *qiniu.Qiniu

func SolveMdFile(filePath string) error {
	// 读取配置文件
	AccessKey := GCF.UString("qiniu.access_key")
	SecretKey := GCF.UString("qiniu.secret_key")
	BucketName := GCF.UString("qiniu.bucket_name")
	Domain := GCF.UString("qiniu.domain")

	GQiniuClient = qiniu.NewQiniuClient(AccessKey, SecretKey, BucketName)

	// 获取文件内容
	file, _ := os.Open(filePath)
	defer file.Close()

	fileContentBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	fileContent := string(fileContentBytes)

	// 匹配所有图片字符
	reg := regexp.MustCompile(`!\[.*\]\(((.*)\.assets/(.*)\.(.*))\)`)
	// [![image-20210808202426899](README.assets/image-20210808202426899.png) README.assets/image-20210808202426899.png README image-20210808202426899 png]
	pics := reg.FindAllStringSubmatch(fileContent, -1)

	// 上传资源图片
	if len(pics) == 0 {
		return nil
	}
	mdName := pics[0][2]
	dirPath := strings.Replace(filePath, mdName+".md", "", -1)
	for _, item := range pics {
		filename := mdName + "-" + item[3]
		NewPicture(GQiniuClient, filename, dirPath+"/"+item[1]).Upload()
		// 替换所有字符串
		fileContent = strings.Replace(fileContent, item[1], Domain+"/"+filename, 1)
	}

	// 回写到新文件夹中
	outputPath := GCF.UString("outdir") + "/" + mdName + ".md"
	if err := ioutil.WriteFile(outputPath, []byte(fileContent), 0666); err != nil {
		fmt.Println(err)
	}

	return nil
}
