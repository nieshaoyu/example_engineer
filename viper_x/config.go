package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Env      string
	Database Database
}

type Database struct {
	Username string `yaml:"Username,omitempty"`
	Password string `yaml:"Password,omitempty"`
	Port     int    `yaml:"Port,omitempty"`
	Drive    string `yaml:"Drive,omitempty"`
}

var C = Config{
	Database: Database{
		Username: "",
		Password: "",
		Port:     0,
		Drive:    "",
	},
	Env: "",
}

func main() {
	parseConfig("./viper_x/config.yaml")
	fmt.Println(C.Database)

}

func parseConfig(path string) {
	logger := log.Default()

	// 配置文件路径
	viper.SetConfigFile(path)
	err := viper.MergeInConfig()
	if err != nil {
		logger.Fatal(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		logger.Println("Config file changed:", in.Name)
		// 配置文件变动, 重新序列化
		_ = viper.UnmarshalExact(&C)
	})
	if err = viper.UnmarshalExact(&C); err != nil {
		logger.Fatal(err)
	}

}
