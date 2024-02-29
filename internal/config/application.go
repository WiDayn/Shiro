package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

// 定义配置结构体
type applicationConfig struct {
	Database databaseConfig `toml:"database"`
}

type databaseConfig struct {
	Server   string
	Port     int
	User     string
	Password string
	Dbname   string
}

var Application applicationConfig

func init() {
	if _, err := toml.DecodeFile("config.toml", &Application); err != nil {
		log.Fatalf("Error loading config.toml: %v", err)
	}
}
