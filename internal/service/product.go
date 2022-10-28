package service

import (
	"context"
	"github.com/radityarestan/ecom-core/internal/entity"
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
	var (
		orm    = p.deps.Database.WithContext(ctx)
		userID = ctx.Value("user_id").(uint)
	)

	var product = &entity.Product{
		Name:       req.Name,
		Detail:     req.Detail,
		Price:      req.Price,
		Stock:      req.Stock,
		CategoryID: req.CategoryID,
		UserID:     userID,
	}

	res, err := p.repo.CreateProduct(orm, product)
	if err != nil {
		p.deps.Logger.Error("failed to create product", err)
		return nil, err
	}

	return &dto.CreateProductResponse{
		ID: res.ID,
	}, nil
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
