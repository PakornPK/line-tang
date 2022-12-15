package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

var port string

func LoadConfig() error {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	setPort(viper.Get("PORT"))
	return err
}

func setPort(p interface{}) {
	port = fmt.Sprintf(":%s", p)
}

func GetPort() string {
	return port
}
