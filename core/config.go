package core

import (
	"cms/global"
	"os"

	_ "embed"
	"github.com/spf13/viper"
)

const (
	deploymentMethodK8s = "k8s"
	localPath           = "./config.yaml"
	configmapPath       = "/data/config.yaml"
)

func InitConfig() *viper.Viper {
	// 部署方式环境变量
	path := getEnv()
	if path == localPath {
		// 默认使用本地配置
		return ViperLocal() // 初始化Viper
	} else {
		// 从 configmap 挂载的文件读取配置
		return ViperLocal(path)

	}
}

func getEnv() string {
	// 部署方式环境变量
	deployMethod := os.Getenv("DEPLOYMENT_METHOD")
	var path string
	if deployMethod == deploymentMethodK8s {
		// 从 configmap 挂载的文件读取配置
		path = configmapPath
	} else {
		// 默认使用本地配置
		path = localPath
	}
	global.EnvPath = path
	return path
}
