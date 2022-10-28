package dto

const (
	CreateProductSuccess  = "Product created successfully"
	SearchProductSuccess  = "Product search success"
	DetailProductSuccess  = "Find product detail success"
	SearchProductError    = "Product search cannot be empty"
	CatalogProductSuccess = "Find product catalog success"
)

type (
	CreateProductRequest struct {
		Name       string `json:"name" validate:"required"`
		Detail     string `json:"detail" validate:"required"`
		Stock      uint   `json:"stock" validate:"required"`
		UserID     uint   `json:"user_id" validate:"required"`
		CategoryID uint   `json:"category_id" validate:"required"`
	}

	CreateProductResponse struct {
		ID uint `json:"id" validate:"required"`
	}

	ProductsResponse struct {
		Products []ProductOverviewResponse `json:"products" validate:"required"`
	}

	ProductOverviewResponse struct {
		ID     uint    `json:"id" validate:"required"`
		Stock  uint    `json:"stock" validate:"required"`
		Sold   uint    `json:"sold" validate:"required"`
		Name   string  `json:"name" validate:"required"`
		Photo  string  `json:"photo" validate:"required"`
		Rating float32 `json:"rating" validate:"required"`
		Price  float64 `json:"price" validate:"required"`
	}

	ProductDetailResponse struct {
		ID           uint    `json:"id" validate:"required"`
		Name         string  `json:"name" validate:"required"`
		Detail       string  `json:"detail" validate:"required"`
		Photo        string  `json:"photo" validate:"required"`
		Price        float64 `json:"price" validate:"required"`
		Stock        uint    `json:"stock" validate:"required"`
		Sold         uint    `json:"sold" validate:"required"`
		Rating       float32 `json:"rating" validate:"required"`
		AmountRating uint    `json:"amount_rating" validate:"required"`

		// TODO: add category + user
	}
)
