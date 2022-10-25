package dto

const (
	// StatusSuccess is a constant for success status
	StatusSuccess = "success"
	// StatusError is a constant for error status
	StatusError = "error"
)

type Response struct {
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
}
