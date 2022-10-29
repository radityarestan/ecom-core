package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/radityarestan/ecom-core/internal/entity"
	"gorm.io/gorm"
	"time"
)

type (
	Product interface {
		CreateProduct(orm *gorm.DB, product *entity.Product) (*entity.Product, error)
		GetBaseProducts(orm *gorm.DB, rc *redis.Client, limit int, offset int) ([]*entity.Product, ResponseType, error)
		FindProducts(orm *gorm.DB, rc *redis.Client, limit int, offset int, search string) ([]*entity.Product, ResponseType, error)
		FindProductByID(orm *gorm.DB, rc *redis.Client, id uint) (*entity.Product, ResponseType, error)

		GetProductRecommendFromRedis(rc *redis.Client, userID uint) ([]*entity.Product, error)
		SaveProductSearchToRedis(rc *redis.Client, search string, value interface{}, exp int) error
		SaveProductDetailToRedis(rc *redis.Client, id uint, value interface{}, exp int) error
		SaveBaseProductToRedis(rc *redis.Client, value interface{}, exp int) error
		SaveProductRecommendToRedis(rc *redis.Client, userID uint, value interface{}, exp int) error
	}

	productRepo struct{}
)

func (p *productRepo) CreateProduct(orm *gorm.DB, product *entity.Product) (*entity.Product, error) {
	if err := orm.Create(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (p *productRepo) GetBaseProducts(orm *gorm.DB, rc *redis.Client, limit int, offset int) ([]*entity.Product, ResponseType, error) {
	var (
		products        []*entity.Product
		productsEncoded []byte
	)

	redisCTX, redisCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer redisCancel()

	err := rc.Get(redisCTX, RedisBaseProductKey).Scan(&productsEncoded)
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, ErrorResponseType, err
	}

	if len(productsEncoded) > 0 {
		if err := json.Unmarshal(productsEncoded, &products); err != nil {
			return nil, ErrorResponseType, err
		}
		return products, RedisResponseType, nil
	}

	if err := orm.Order("updated_at DESC").Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, ErrorResponseType, err
	}

	return products, PostgresResponseType, nil
}

func (p *productRepo) FindProducts(orm *gorm.DB, rc *redis.Client, limit int, offset int, search string) ([]*entity.Product, ResponseType, error) {
	var (
		products        []*entity.Product
		productsEncoded []byte
	)

	redisCTX, redisCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer redisCancel()

	err := rc.Get(redisCTX, fmt.Sprintf(RedisProductSearchKey, search)).Scan(&productsEncoded)
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, ErrorResponseType, err
	}

	if len(productsEncoded) > 0 {
		if err := json.Unmarshal(productsEncoded, &products); err != nil {
			return nil, ErrorResponseType, err
		}
		return products, RedisResponseType, nil
	}

	if err := orm.Where("LOWER(name) LIKE ?", "%"+search+"%").Order("updated_at DESC").Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, ErrorResponseType, err
	}

	return products, PostgresResponseType, nil
}

func (p *productRepo) FindProductByID(orm *gorm.DB, rc *redis.Client, id uint) (*entity.Product, ResponseType, error) {
	var (
		product        = &entity.Product{}
		productEncoded []byte
	)

	redisCTX, redisCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer redisCancel()

	err := rc.Get(redisCTX, fmt.Sprintf(RedisProductDetailKey, id)).Scan(&productEncoded)
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, ErrorResponseType, err
	}

	if len(productEncoded) > 0 {
		if err := json.Unmarshal(productEncoded, &product); err != nil {
			return nil, ErrorResponseType, err
		}
		return product, RedisResponseType, nil
	}

	if err := orm.Where("id = ?", id).First(product).Error; err != nil {
		return nil, ErrorResponseType, err
	}

	return product, PostgresResponseType, nil
}

func (p *productRepo) GetProductRecommendFromRedis(rc *redis.Client, userID uint) ([]*entity.Product, error) {
	var (
		products        []*entity.Product
		productsEncoded []byte
	)

	redisCTX, redisCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer redisCancel()

	err := rc.Get(redisCTX, fmt.Sprintf(RedisProductRecommendKey, userID)).Scan(&productsEncoded)
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}

	if len(productsEncoded) == 0 {
		return products, nil
	}

	if err := json.Unmarshal(productsEncoded, &products); err != nil {
		return nil, err
	}

	return products, nil
}

func (p *productRepo) SaveProductSearchToRedis(rc *redis.Client, search string, value interface{}, exp int) error {
	encodedValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	redisCTX, redisCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer redisCancel()

	op := rc.Set(redisCTX, fmt.Sprintf(RedisProductSearchKey, search), encodedValue, time.Duration(exp)*time.Minute)
	if err := op.Err(); err != nil {
		return err
	}

	return nil
}

func (p *productRepo) SaveProductDetailToRedis(rc *redis.Client, id uint, value interface{}, exp int) error {
	encodedValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	redisCTX, redisCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer redisCancel()

	op := rc.Set(redisCTX, fmt.Sprintf(RedisProductDetailKey, id), encodedValue, time.Duration(exp)*time.Minute)
	if err := op.Err(); err != nil {
		return err
	}

	return nil
}

func (p *productRepo) SaveBaseProductToRedis(rc *redis.Client, value interface{}, exp int) error {
	encodedValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	redisCTX, redisCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer redisCancel()

	op := rc.Set(redisCTX, RedisBaseProductKey, encodedValue, time.Duration(exp)*time.Minute)
	if err := op.Err(); err != nil {
		return err
	}

	return nil
}

func (p *productRepo) SaveProductRecommendToRedis(rc *redis.Client, userID uint, value interface{}, exp int) error {
	products, err := p.GetProductRecommendFromRedis(rc, userID)
	if err != nil {
		return err
	}

	productToSave := value.(*entity.Product)
	for _, product := range products {
		if product.ID == productToSave.ID {
			return nil
		}
	}

	products = append(products, productToSave)
	encodedValue, err := json.Marshal(products)
	if err != nil {
		return err
	}

	redisCTX, redisCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer redisCancel()

	op := rc.Set(redisCTX, fmt.Sprintf(RedisProductRecommendKey, userID), encodedValue, time.Duration(exp)*time.Minute)
	if err := op.Err(); err != nil {
		return err
	}

	return nil
}

func NewProduct() (Product, error) {
	return &productRepo{}, nil
}
