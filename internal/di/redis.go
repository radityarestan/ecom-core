package di

import (
	"github.com/go-redis/redis/v8"
	"github.com/radityarestan/ecom-core/internal/shared/config"
)

func NewRedis(config *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host + ":" + config.Redis.Port,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})
}
