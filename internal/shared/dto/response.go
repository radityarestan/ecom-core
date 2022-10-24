package dto

const (
	// StatusSuccess is a constant for success status
	StatusSuccess = "Success"
	// StatusError is a constant for error status
	StatusError = "Error"
)

type Response struct {
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
}
