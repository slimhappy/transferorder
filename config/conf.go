package config

import (
	"github.com/spf13/viper"
)

type ConfYaml struct {
	Svr SectionSvr `yaml:"svr"`
}

type SectionSvr struct {
	Port    string `yaml:"port"`
	FileDir string `yaml:"file_dir"`
}

var Conf ConfYaml

func init() {
	// 设定读取配置的路径
	viper.AddConfigPath("./config/")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(err)
	}
	Conf.Svr.Port = viper.GetString("svr.port")
	Conf.Svr.FileDir = viper.GetString("svr.file_dir")
}
