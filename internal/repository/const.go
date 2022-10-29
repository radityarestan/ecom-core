package repository

const (
	RedisBaseProductKey      = "product_base"
	RedisProductRecommendKey = "product_recommend#%d"
	RedisProductSearchKey    = "product_search#%s"
	RedisProductDetailKey    = "product_detail#%d"
)

type ResponseType string

var (
	RedisResponseType    = ResponseType("Redis")
	PostgresResponseType = ResponseType("Postgres")
	ErrorResponseType    = ResponseType("Error")
)

func (r *ResponseType) String() string {
	return string(*r)
}
