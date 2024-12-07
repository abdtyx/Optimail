package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	JWTKey    string           `mapstructure:"jwtkey"`
	MicroUser *ConfigMicroUser `mapstructure:"microuser"`
	ChatGPT   *ConfigGPT       `mapstructure:"gpt"`
	Webpage   *ConfigWebpage   `mapstructure:"webpage"`
}

type ConfigGPT struct {
	Apiurl string `mapstructure:"apiurl"`
	Model  string `mapstructure:"model"`
	Apikey string `mapstructure:"apikey"`
}

type ConfigMicroUser struct {
	GrpcAddr string `mapstructure:"grpc_addr"`
}

type ConfigWebpage struct {
	BasePath string `mapstructure:"base_path"`
}

func New() *Config {
	return &Config{
		ChatGPT: &ConfigGPT{
			Apiurl: "https://api.openai.com/v1/chat/completions",
			Model:  "gpt-3.5-turbo",
			Apikey: "",
		},
		MicroUser: &ConfigMicroUser{
			GrpcAddr: "localhost:50051",
		},
		Webpage: &ConfigWebpage{
			BasePath: "../",
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
	if tmp.JWTKey != "" {
		cfg.JWTKey = tmp.JWTKey
	} else {
		panic("JWTKey required in config.toml!")
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
	if tmp.MicroUser != nil {
		if tmp.MicroUser.GrpcAddr != "" {
			cfg.MicroUser.GrpcAddr = tmp.MicroUser.GrpcAddr
		}
	}
	if tmp.Webpage != nil {
		if tmp.Webpage.BasePath != "" {
			cfg.Webpage.BasePath = tmp.Webpage.BasePath
		}
	}
}

func (cfg *Config) String() string {
	return cfg.JWTKey + "\n" + cfg.ChatGPT.Apiurl + "\n" + cfg.ChatGPT.Model + "\n" + cfg.ChatGPT.Apikey + "\n" + cfg.MicroUser.GrpcAddr + "\n" + cfg.Webpage.BasePath + "\n"
}
