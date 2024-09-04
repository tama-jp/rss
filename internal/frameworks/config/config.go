package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type Config struct {
	DB      DBConfig      `toml:"database"`
	Admin   AdminConfig   `toml:"admin"`
	General GeneralConfig `toml:"general"`
	Logger  LoggerConfig  `toml:"logger"`
	Rooting RootingConfig `toml:"rooting"`
}

type DBConfig struct {
	DataBase string `toml:"data_base"`
	FileName string `toml:"file_name"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Name     string `toml:"name"`
	Debug    int    `toml:"debug"`
}

type RootingConfig struct {
	Port int `toml:"port"`
}

type AdminConfig struct {
	UserName  string `toml:"user_name"`
	LastName  string `toml:"last_name"`
	FirstName string `toml:"first_name"`
	Password  string `toml:"password"`
}

type GeneralConfig struct {
	Debug       int    `toml:"debug"`
	LogFileName string `toml:"log_file_name"`
}

type LoggerConfig struct {
	FileName   string `toml:"file_name"`
	MaxSize    int    `toml:"max_size"`
	MaxBackups int    `toml:"max_backups"`
	MaxAge     int    `toml:"max_age"`
}

func NewConfig() (*Config, error) {

	var conf Config
	var confPath string

	appMode := os.Getenv("APP_MODE")
	if appMode == "" {
		appMode = "config"
	}

	confPath = appMode + ".toml"

	fmt.Println("confPath:" + confPath)

	if _, err := toml.DecodeFile(confPath, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
