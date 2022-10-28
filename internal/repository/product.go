package repository

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/radityarestan/ecom-core/internal/entity"
	"gorm.io/gorm"
	"strings"
	"time"
)

type (
	Product interface {
		CreateProduct(orm *gorm.DB, product *entity.Product) (*entity.Product, error)
		GetBaseProducts(orm *gorm.DB, rc *redis.Client, limit int, offset int) ([]*entity.Product, error)
		FindProducts(orm *gorm.DB, rc *redis.Client, limit int, offset int, search string) ([]*entity.Product, error)
		FindProductByID(orm *gorm.DB, rc *redis.Client, id uint) (*entity.Product, error)

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

func (p *productRepo) GetBaseProducts(orm *gorm.DB, rc *redis.Client, limit int, offset int) ([]*entity.Product, error) {
	var products []*entity.Product

	if err := rc.Get(redisCTX, RedisBaseProductKey).Scan(&products); err != nil {
		return nil, err
	}

	if len(products) > 0 {
		return products, nil
	}

	if err := orm.Limit(limit).Offset(offset).Order("updated_at DESC").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (p *productRepo) FindProducts(orm *gorm.DB, rc *redis.Client, limit int, offset int, search string) ([]*entity.Product, error) {
	var products []*entity.Product

	if err := rc.Get(redisCTX, fmt.Sprintf(RedisProductSearchKey, strings.ToLower(search))).Scan(&products); err != nil {
		return nil, err
	}

	if len(products) > 0 {
		return products, nil
	}

	if err := orm.Where("name LIKE ?", "%"+search+"%").Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (p *productRepo) FindProductByID(orm *gorm.DB, rc *redis.Client, id uint) (*entity.Product, error) {
	var product = &entity.Product{}

	if err := rc.Get(redisCTX, fmt.Sprintf(RedisProductDetailKey, id)).Scan(&product); err != nil {
		return nil, err
	}

	if product.ID > 0 {
		return product, nil
	}

	if err := orm.Where("id = ?", id).First(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (p *productRepo) GetProductRecommendFromRedis(rc *redis.Client, userID uint) ([]*entity.Product, error) {
	var products []*entity.Product

	key := fmt.Sprintf(RedisProductRecommendKey, userID)
	if err := rc.Get(redisCTX, key).Scan(&products); err != nil {
		return nil, err
	}

	return products, nil
}

func (p *productRepo) SaveProductSearchToRedis(rc *redis.Client, search string, value interface{}, exp int) error {
	defer redisCancel()

	op := rc.Set(redisCTX, fmt.Sprintf(RedisProductSearchKey, strings.ToLower(search)), value, time.Duration(exp)*time.Hour)
	if err := op.Err(); err != nil {
		return err
	}

	return nil
}

func (p *productRepo) SaveProductDetailToRedis(rc *redis.Client, id uint, value interface{}, exp int) error {
	defer redisCancel()

	op := rc.Set(redisCTX, fmt.Sprintf(RedisProductDetailKey, id), value, time.Duration(exp)*time.Hour)
	if err := op.Err(); err != nil {
		return err
	}

	return nil
}

func (p *productRepo) SaveBaseProductToRedis(rc *redis.Client, value interface{}, exp int) error {
	defer redisCancel()

	op := rc.Set(redisCTX, RedisBaseProductKey, value, time.Duration(exp)*time.Hour)
	if err := op.Err(); err != nil {
		return err
	}

	return nil
}

func (p *productRepo) SaveProductRecommendToRedis(rc *redis.Client, userID uint, value interface{}, exp int) error {
	defer redisCancel()

	op := rc.Set(redisCTX, fmt.Sprintf(RedisProductRecommendKey, userID), value, time.Duration(exp)*time.Hour)
	if err := op.Err(); err != nil {
		return err
	}

	return nil
}

func NewProduct() (Product, error) {
	return &productRepo{}, nil
}
