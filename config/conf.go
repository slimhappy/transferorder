package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type ConfYaml struct {
	Svr SectionSvr `yaml:"svr"`
}

type SectionSvr struct {
	Port string `yaml:"port"`
}

var Conf ConfYaml

func init() {
	// 设定读取配置的路径
	viper.AddConfigPath("./config/")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	Conf.Svr.Port = viper.GetString("svr.port")
}
