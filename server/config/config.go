package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DSN     string     `mapstructure:"dsn"`
	ChatGPT *ConfigGPT `mapstructure:"gpt"`
}

type ConfigGPT struct {
	Apiurl string
	Model  string
	Apikey string
}

func New() *Config {
	return &Config{
		DSN: "root:password@tcp(localhost:3306)/test",
		ChatGPT: &ConfigGPT{
			Apiurl: "https://api.openai.com/v1/chat/completions",
			Model:  "gpt-3.5-turbo",
			Apikey: "",
		},
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
	if tmp.ChatGPT != nil {
		if tmp.ChatGPT.Apikey != "" {
			cfg.ChatGPT.Apikey = tmp.ChatGPT.Apikey
		} else {
			panic("Please set the API_KEY from OPENAI")
		}
		if tmp.ChatGPT.Apiurl != "" {
			cfg.ChatGPT.Apiurl = tmp.ChatGPT.Apiurl
		}
		if tmp.ChatGPT.Model != "" {
			cfg.ChatGPT.Model = tmp.ChatGPT.Model
		}
	}
}
