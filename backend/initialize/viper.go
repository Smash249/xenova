package initialize

import (
	"fmt"

	"github.com/spf13/viper"
)

func initViper() {
	viper.SetConfigName("system")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}
}
