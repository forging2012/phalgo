package phalgo

//	PhalGo-Config
//	使用spf13大神的viper配置文件获取工具作为phalgo的配置文件工具
//	喵了个咪 <wenzhenxi@vip.qq.com> 2016/5/11
//  依赖情况:
//          "github.com/spf13/viper"

import (
	"github.com/spf13/viper"
	"fmt"
)


var Config *viper.Viper

//初始化配置文件
func NewConfig(filePath string, fileName string) {
	Config = viper.New()
	Config.SetConfigName(fileName)
	Config.AddConfigPath(filePath + "/")

	err := Config.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err).Error())
	}
}