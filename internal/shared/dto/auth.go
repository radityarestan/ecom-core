package dto

const (
	CreateUserSuccess = "User created successfully"
)

type (
	TestRequest struct {
		Name string `json:"name" validate:"required"`
	}

	TestResponse struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
)
