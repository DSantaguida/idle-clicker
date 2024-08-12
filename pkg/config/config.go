package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig;
	Db DBConfig;
}

type ServerConfig struct {
	Version string;
	Port string;
}

type DBConfig struct {
	Host string;
	Port string;
	User string;
	Password string;
	Name string;
	Driver string;
	SslMode string;
}

func loadConfig(path string, name string) (*viper.Viper, error){
	v := viper.New()

	v.SetConfigName(name)
	v.SetConfigType("yml")
	v.AddConfigPath(path)

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return v, nil
}

func parseConfig(v *viper.Viper) (*Config, error){
	var config Config
	
	err := v.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func GetConfig(path string, name string) (*Config, error){
	v, err := loadConfig(path, name)
	if err != nil {
		return nil, err
	}

	config, err := parseConfig(v)
	if err != nil {
		return nil, err
	}

	return config, nil
}
