package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	MySQL  MySQLConfig
	Redis  RedisConfig
}

// ServerConfig Server相关的配置
type ServerConfig struct {
	Port string
}

// MySQLConfig MySQL相关的配置
type MySQLConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// RedisConfig Redis相关的配置
type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

var AppConfig *Config

// InitConfig 利用Viper读取配置文件
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	// 是否读取成功
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	AppConfig = &Config{}

	// 将读取的配置转换为go的结构体，也就是Config结构体
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}
}
