package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerConfig ServerConfig
	JWTConfig    JWTConfig
	DBConfig     DBConfig
	Logger       Logger
}

type ServerConfig struct {
	AppVersion       string
	Mode             string
	RoutePrefix      string
	Debug            bool
	Port             string
	TimeoutSecs      int64
	ReadTimeoutSecs  int64
	WriteTimeoutSecs int64
}

type JWTConfig struct {
	SessionTime int
	SecretKey   string
}

type DBConfig struct {
	DataSourceName  string
	Name            string
	MigrationFolder string
	MaxOpen         int
	MaxIdle         int
	MaxLifetime     int
}

// Logger config
type Logger struct {
	Development bool
	Encoding    string
	Level       string
}

// LoadConfig file from given path
func LoadConfig(filename string) (*Config, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil

}
