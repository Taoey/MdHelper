package main

import (
	"fmt"

	"github.com/Taoey/MdHelper/service"
)

func main() {

	// 初始化配置
	if err := service.InitConfig(); err != nil {
		panic(err)
	}

	fmt.Println(service.GCF)

	// 获取文件路径
	for {
		var mdPath string
		fmt.Println("请输入文件路径:")
		fmt.Scanf("%s", &mdPath)

		// 进行转化
		service.SolveMdFile(mdPath)
	}
}
