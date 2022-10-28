package repository

import (
	"github.com/radityarestan/ecom-core/internal/entity"
	"gorm.io/gorm"
)

type (
	Product interface {
		CreateProduct(orm *gorm.DB, product *entity.Product) (*entity.Product, error)
		GetBaseProducts(orm *gorm.DB, limit int) ([]*entity.Product, error)
		FindProducts(orm *gorm.DB, limit int, offset int, search string) ([]*entity.Product, error)
		FindProductByID(orm *gorm.DB, id int64) (*entity.Product, error)
	}

	productRepo struct{}
)

func (p *productRepo) CreateProduct(orm *gorm.DB, product *entity.Product) (*entity.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p *productRepo) GetBaseProducts(orm *gorm.DB, limit int) ([]*entity.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p *productRepo) FindProducts(orm *gorm.DB, limit int, offset int, search string) ([]*entity.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p *productRepo) FindProductByID(orm *gorm.DB, id int64) (*entity.Product, error) {
	//TODO implement me
	panic("implement me")
}

func NewProduct() (Product, error) {
	return &productRepo{}, nil
}
