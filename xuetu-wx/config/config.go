package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}

	Redis struct {
		Addr     string
		DB       int
		Password string
	}
}

var AppConfig *Config

// InitConfig 利用Viper读取配置文件
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	AppConfig = &Config{}
	// 将读取的配置转换为go的结构体，也就是Config结构体
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	// 实例化redis
	InitRedis()
}
