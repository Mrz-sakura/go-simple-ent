package config

import (
	"fmt"

	"dwbackend/global"

	"github.com/spf13/viper"
)

func InitGlobalConfig() {
	global.GlobalConfigInstance = &global.GlobalConfig{}

	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("cfg")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(global.GlobalConfigInstance)
	if err != nil {
		panic(err)
	}

	fmt.Println("配置加载成功～")
}
