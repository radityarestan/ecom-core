package service

import (
	"context"
	"github.com/radityarestan/ecom-core/internal/repository"
	"github.com/radityarestan/ecom-core/internal/shared"
	"github.com/radityarestan/ecom-core/internal/shared/dto"
)

type (
	Product interface {
		CreateProduct(ctx context.Context, req *dto.CreateProductRequest) (*dto.CreateProductResponse, error)
		GetBaseProducts(ctx context.Context, userID uint) (*dto.ProductsResponse, error)
		FindProducts(ctx context.Context, search string) (*dto.ProductsResponse, error)
		FindProductByID(ctx context.Context, id uint) (*dto.ProductDetailResponse, error)
	}

	productService struct {
		deps shared.Deps
		repo repository.Product
	}
)

func (p *productService) CreateProduct(ctx context.Context, req *dto.CreateProductRequest) (*dto.CreateProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p *productService) GetBaseProducts(ctx context.Context, userID uint) (*dto.ProductsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p *productService) FindProducts(ctx context.Context, search string) (*dto.ProductsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p *productService) FindProductByID(ctx context.Context, id uint) (*dto.ProductDetailResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewProduct(deps shared.Deps, repo repository.Product) (Product, error) {
	return &productService{
		deps: deps,
		repo: repo,
	}, nil
}
