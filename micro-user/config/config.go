package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DSN      string `mapstructure:"dsn"`
	GrpcAddr string `mapstructure:"grpc_addr"`
}

func New() *Config {
	return &Config{
		DSN:      "root:password@tcp(localhost:3306)/test",
		GrpcAddr: "localhost:50051",
	}
}

func (cfg *Config) Load() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic("failed to read config file, error: " + err.Error())
	}
	tmp := &Config{}
	if err := viper.UnmarshalKey("server", tmp); err != nil {
		panic("failed to init server config, error: " + err.Error())
	}
	if tmp.DSN != "" {
		cfg.DSN = tmp.DSN
	}
	if tmp.GrpcAddr != "" {
		cfg.GrpcAddr = tmp.GrpcAddr
	}
}
