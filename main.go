package main

import (
	"flag"
	"fmt"

	"cms/common/enums"
	"cms/core"
	"cms/global"

)


func main() {
	// -env=xxx 指定启动环境
	// -env=local 本地; -env=dev 开发环境; -env=test2 测试环境
	// -env=prod -c=config.prod.yaml
	envFlag := flag.String("env", enums.LOCAL.Code, enums.ShowEnvs())
	flag.Parse()

	global.Env = enums.ToEnv(*envFlag)
	fmt.Println("command param env=========", global.Env.String())

	global.GVA_VP = core.InitConfig()
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	systemEnv, ok := global.GVA_VP.Get("system.env").(string)
	if !ok {
		fmt.Println("has no system.env")
	}
	fmt.Println("command param env=========", systemEnv)
	if systemEnv == "local" {
		global.Env = enums.LOCAL
	} else if systemEnv == "dev" {
		global.Env = enums.DEV
	} else if systemEnv == "prod" {
		global.Env = enums.PROD
	} else {
		global.Env = enums.LOCAL
	}
	 // gorm连接数据库
	// global.GVA_DB = initialize.Gorm()

	core.RunWindowsServer()
}
