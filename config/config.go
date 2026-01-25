package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	MySQL    MySQLConfig
	Redis    RedisConfig
	RabbitMQ RabbitMQConfig
	JWT      JWTConfig
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
	Host     string
	Port     int
	Password string
	DB       int
}

// RabbitMQConfig RabbitMQ相关的配置
type RabbitMQConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// JWTConfig JWT相关的配置
type JWTConfig struct {
	Secret string `destructure:"secret"`
	Expire int64  `destructure:"expire"` // 过期时间，单位秒
}

// AppConfig 全局配置变量
var AppConfig *Config

// InitConfig 利用Viper读取配置文件
func InitConfig() {
	// 设置配置文件名、类型、路径
	viper.SetConfigName("config")   // 指定配置文件名称（不带后缀）
	viper.SetConfigType("yml")      // 指定配置文件类型
	viper.AddConfigPath("./config") // 指定配置文件路径

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	// 实例化配置结构体,即初始化 AppConfig
	AppConfig = &Config{}

	// 将读取的配置转换为go的结构体，也就是Config结构体
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}
}
