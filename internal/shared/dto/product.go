package dto

const (
	CreateProductSuccess  = "Product created successfully"
	SearchProductSuccess  = "Product search success"
	DetailProductSuccess  = "Find product detail success"
	CatalogProductSuccess = "Find product catalog success"
)

type (
	CreateProductRequest struct {
		Name       string  `json:"name" validate:"required"`
		Detail     string  `json:"detail" validate:"required"`
		Price      float64 `json:"price" validate:"required"`
		Stock      uint    `json:"stock" validate:"required"`
		CategoryID uint    `json:"category_id" validate:"required,gte=1"`
	}

	CreateProductResponse struct {
		ID uint `json:"id" validate:"required"`
	}

	ProductsResponse struct {
		RecommendedProducts []ProductOverviewResponse `json:"products_history" validate:"required"`
		Products            []ProductOverviewResponse `json:"products" validate:"required"`
	}

	ProductOverviewResponse struct {
		ID     uint    `json:"id" validate:"required"`
		Stock  uint    `json:"stock" validate:"gte=0"`
		Sold   uint    `json:"sold" validate:"gte=0"`
		Name   string  `json:"name" validate:"required"`
		Photo  string  `json:"photo" validate:"required"`
		Rating float32 `json:"rating" validate:"gte=0,lte=5"`
		Price  float64 `json:"price" validate:"required"`
	}

	ProductDetailResponse struct {
		ID           uint    `json:"id" validate:"required"`
		Stock        uint    `json:"stock" validate:"gte=0"`
		Sold         uint    `json:"sold" validate:"gte=0"`
		AmountRating uint    `json:"amount_rating" validate:"gte=0"`
		Name         string  `json:"name" validate:"required"`
		Detail       string  `json:"detail" validate:"required"`
		Photo        string  `json:"photo" validate:"required"`
		Rating       float32 `json:"rating" validate:"gte=0,lte=5"`
		Price        float64 `json:"price" validate:"required"`
	}
)
