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

type RedisConfig struct {
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	Password string `mapstructure:"PASSWORD"`
	DB       int    `mapstructure:"DB"`
}

type NSQConfig struct {
	Host    string `mapstructure:"HOST"`
	Port    string `mapstructure:"PORT"`
	Topic   string `mapstructure:"TOPIC"`
	Channel string `mapstructure:"CHANNEL"`
}

type GCPConfig struct {
	Credential         string `mapstructure:"CREDENTIAL"`
	ProjectID          string `mapstructure:"PROJECT_ID"`
	BucketName         string `mapstructure:"BUCKET_NAME"`
	ProductStoragePath string `mapstructure:"PRODUCT_STORAGE_PATH"`
}

type Config struct {
	Server   ServerConfig   `mapstructure:"SERVER"`
	Database DatabaseConfig `mapstructure:"DATABASE"`
	Redis    RedisConfig    `mapstructure:"REDIS"`
	NSQ      NSQConfig      `mapstructure:"NSQ"`
	GCP      GCPConfig      `mapstructure:"GCP"`
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
