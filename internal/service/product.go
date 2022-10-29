package service

import (
	"context"
	"github.com/radityarestan/ecom-core/internal/entity"
	"github.com/radityarestan/ecom-core/internal/repository"
	"github.com/radityarestan/ecom-core/internal/shared"
	"github.com/radityarestan/ecom-core/internal/shared/dto"
	"strings"
)

type (
	Product interface {
		CreateProduct(ctx context.Context, req *dto.CreateProductRequest) (*dto.CreateProductResponse, error)
		GetProductCatalog(ctx context.Context, userID uint, limit int, offset int) (*dto.ProductsResponse, error)
		FindProducts(ctx context.Context, search string, limit int, offset int) ([]*dto.ProductOverviewResponse, error)
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
		Photo:      "default.jpeg",
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

func (p *productService) GetProductCatalog(ctx context.Context, userID uint, limit int, offset int) (*dto.ProductsResponse, error) {
	var (
		orm = p.deps.Database.WithContext(ctx)
		rc  = p.deps.Redis.WithContext(ctx)
	)

	recommendedProducts, err := p.repo.GetProductRecommendFromRedis(rc, userID)
	if err != nil {
		p.deps.Logger.Error("failed to get recommended products from redis", err)
		return nil, err
	}

	products, respType, err := p.repo.GetBaseProducts(orm, rc, limit, offset)
	if err != nil {
		p.deps.Logger.Error("failed to get base products", err)
		return nil, err
	}

	var res = &dto.ProductsResponse{
		RecommendedProducts: make([]dto.ProductOverviewResponse, 0),
		Products:            make([]dto.ProductOverviewResponse, 0),
	}

	for _, product := range recommendedProducts {
		res.RecommendedProducts = append(res.RecommendedProducts, dto.ProductOverviewResponse{
			ID:     product.ID,
			Stock:  product.Stock,
			Sold:   product.Sold,
			Name:   product.Name,
			Photo:  product.Photo,
			Rating: product.Rating,
			Price:  product.Price,
		})
	}

	for _, product := range products {
		res.Products = append(res.Products, dto.ProductOverviewResponse{
			ID:     product.ID,
			Stock:  product.Stock,
			Sold:   product.Sold,
			Name:   product.Name,
			Photo:  product.Photo,
			Rating: product.Rating,
			Price:  product.Price,
		})
	}

	go func() {
		if respType == repository.RedisResponseType {
			return
		}

		p.deps.Logger.Info("saving product catalog to redis")

		err := p.repo.SaveBaseProductToRedis(rc, products, 5)
		if err != nil {
			p.deps.Logger.Error("failed to save product recommend to redis", err)
		}
	}()

	return res, nil

}

func (p *productService) FindProducts(ctx context.Context, search string, limit int, offset int) ([]*dto.ProductOverviewResponse, error) {
	var (
		orm = p.deps.Database.WithContext(ctx)
		rc  = p.deps.Redis.WithContext(ctx)
	)

	products, respType, err := p.repo.FindProducts(orm, rc, limit, offset, strings.ToLower(search))
	if err != nil {
		p.deps.Logger.Error("failed to find products", err)
		return nil, err
	}

	var res = make([]*dto.ProductOverviewResponse, 0)
	for _, product := range products {
		res = append(res, &dto.ProductOverviewResponse{
			ID:     product.ID,
			Stock:  product.Stock,
			Sold:   product.Sold,
			Name:   product.Name,
			Photo:  product.Photo,
			Rating: product.Rating,
			Price:  product.Price,
		})
	}

	go func() {
		if respType == repository.RedisResponseType {
			return
		}

		p.deps.Logger.Info("saving product search to redis")

		err := p.repo.SaveProductSearchToRedis(rc, strings.ToLower(search), products, 5)
		if err != nil {
			p.deps.Logger.Error("failed to save product recommend to redis", err)
		}
	}()

	return res, nil
}

func (p *productService) FindProductByID(ctx context.Context, id uint) (*dto.ProductDetailResponse, error) {
	var (
		orm    = p.deps.Database.WithContext(ctx)
		rc     = p.deps.Redis.WithContext(ctx)
		userID = ctx.Value("user_id").(uint)
	)

	product, respType, err := p.repo.FindProductByID(orm, rc, id)
	if err != nil {
		p.deps.Logger.Error("failed to find product by id", err)
		return nil, err
	}

	go func() {
		if respType != repository.RedisResponseType {
			p.deps.Logger.Info("saving product detail to redis")

			err := p.repo.SaveProductDetailToRedis(rc, id, product, 60)
			if err != nil {
				p.deps.Logger.Error("failed to save product detail to redis", err)
			}
		}

		err = p.repo.SaveProductRecommendToRedis(rc, userID, product, 180)
		if err != nil {
			p.deps.Logger.Error("failed to save product recommend to redis", err)
		}
	}()

	return &dto.ProductDetailResponse{
		ID:           product.ID,
		Name:         product.Name,
		Detail:       product.Detail,
		Photo:        product.Photo,
		Price:        product.Price,
		Stock:        product.Stock,
		Sold:         product.Sold,
		Rating:       product.Rating,
		AmountRating: product.AmountRating,
	}, nil
}

func NewProduct(deps shared.Deps, repo repository.Product) (Product, error) {
	return &productService{
		deps: deps,
		repo: repo,
	}, nil
}
