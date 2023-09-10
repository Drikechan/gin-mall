package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var Config *Conf

type System struct {
	Port string `yaml:"port"`
}

type Mysql struct {
	Dialect  string `yaml:"dialect"`
	DbHost   string `yaml:"dbHost"`
	DbPort   string `yaml:"dbPort"`
	DbName   string `yaml:"dbName"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

type Es struct {
	EsHost  string `yaml:"esHost"`
	EsPort  string `yaml:"esPort"`
	EsIndex string `yaml:"esIndex"`
}

type Conf struct {
	System *System           `yaml:"system"`
	Mysql  map[string]*Mysql `yaml:"mysql"`
	Es     *Es               `yaml:"es"`
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config/locales")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("读取文件失败")
		return
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		fmt.Println("写入配置序列化错误！")
		return
	}
}
