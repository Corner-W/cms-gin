package core

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/songzhibin97/gkit/cache/local_cache"

	"cms/global"
	"cms/utils"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func ViperLocal(path ...string) *viper.Viper {
	//配置文件名
	var config string
	// 优先级: 命令行 > 环境变量 > 配置文件
	if len(path) == 0 {
		//命令行
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" {
			//环境变量
			if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {
				//配置文件
				config = utils.ConfigFile
				fmt.Printf("您正在使用config文件,config的路径为%v\n", config)
			} else {
				config = configEnv
				fmt.Printf("您正在使用GVA_CONFIG环境变量,config的路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	//设置配置文件名
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	//读取配置文件中的内容
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}
	// root 适配性
	// 根据root位置去找到对应迁移位置,保证root路径有效
	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(time.Second * time.Duration(global.GVA_CONFIG.JWT.ExpiresTime)),
	)
	return v
}
