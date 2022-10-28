package repository

import (
	"context"
	"time"
)

var (
	redisCTX, redisCancel = context.WithTimeout(context.Background(), 5*time.Second)
)

const (
	RedisBaseProductKey      = "product_base"
	RedisProductRecommendKey = "product_recommend#%d"
	RedisProductSearchKey    = "product_search#%s"
	RedisProductDetailKey    = "product_detail#%d"
)
