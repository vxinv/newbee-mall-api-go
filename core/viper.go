package core

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"main/global"
	"main/utils"
)

func GetConfig(path ...string) string {
	var config string
	if len(path) != 0 {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
		return config
	}
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()
	if config != "" { // 优先级: 命令行 > 环境变量 > 默认值
		fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		return config
	}
	config = utils.ConfigFile
	fmt.Printf("您正在使用config的默认值,config的路径为%v\n", utils.ConfigFile)
	return config
}

func Viper(path ...string) *viper.Viper {
	var config = GetConfig(path...)
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
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
	return v
}
