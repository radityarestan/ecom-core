package config

import "github.com/spf13/viper"

type ServerConfig struct {
	Port string `mapstructure:"PORT"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	DbName   string `mapstructure:"NAME"`
	SSLMode  string `mapstructure:"SSLMODE"`
	Timezone string `mapstructure:"TIMEZONE"`
}

type NSQConfig struct {
	Host    string `mapstructure:"HOST"`
	Port    string `mapstructure:"PORT"`
	Topic   string `mapstructure:"TOPIC"`
	Channel string `mapstructure:"CHANNEL"`
}

type Config struct {
	Server   ServerConfig   `mapstructure:"SERVER"`
	Database DatabaseConfig `mapstructure:"DATABASE"`
	NSQ      NSQConfig      `mapstructure:"NSQ"`
}

func NewConfig() (*Config, error) {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.AutomaticEnv()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
