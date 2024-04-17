package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	//InitConfig在哪执行 这个路径就是哪！！！！！！
	//go_test用
	viper.AddConfigPath("..")
	//main用
	viper.AddConfigPath("./app/")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("[InitConfig]config read fault -- err: %s\n", err.Error()))
		return
	}
}
