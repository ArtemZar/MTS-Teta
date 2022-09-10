package config

import "github.com/spf13/viper"

type Config struct {
	SrvConfig   ServerConfig
	Credentials map[string]string
	SignKey     string
}

type ServerConfig struct {
	Addr string
}

func New() (*Config, error) {
	viper.SetDefault("server.addr", ":8080")
	//viper.SetDefault("credentials", map[string]string{})

	conf := &Config{}
	conf.SrvConfig.Addr = viper.GetString("server.addr")
	conf.Credentials = viper.GetStringMapString("credentials")
	conf.SignKey = viper.GetString("sign_key")

	return conf, nil
}
