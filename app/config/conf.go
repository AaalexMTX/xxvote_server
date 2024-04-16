package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./app/")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("[InitConfig]config read fault -- err: %s\n", err.Error()))
	}
}
