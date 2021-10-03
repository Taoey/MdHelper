package service

import (
	"github.com/Taoey/MdHelper/util"
	"github.com/olebedev/config"
)

var GCF *config.Config //global config
var configFile = "configs/config.yml"

func InitConfig() error {
	var err error
	// 通过绝对路径获取配置文件
	configAbDir := util.GetCurrentAbPath()
	//configAbDir, err := filepath.Abs(".")

	configPath := configAbDir + "/" + configFile

	GCF, err = config.ParseYamlFile(configPath)
	if err != nil {
		return err
	}
	return nil
}
