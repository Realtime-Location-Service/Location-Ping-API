package error

import "net/http"

// Err provides the flexibilty to structure the error
// can be used for err error code
type Err struct {
	StatusCode int    `json:"statusCode"`
	StatusText string `json:"statusText"`
	Message    string `json:"message"`
}

// NewErr returns a new error
func NewErr(statusCode int, errMsg string) *Err {
	return &Err{
		StatusCode: statusCode,
		StatusText: http.StatusText(statusCode),
		Message:    errMsg,
	}
}

const (
	// ErrNoDataFound is returned when data not found for a request.
	ErrNoDataFound = "No data found for the request"
)
