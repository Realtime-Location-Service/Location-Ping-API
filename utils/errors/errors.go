package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

// Err provides the flexibilty to structure the error
// can be used for error code
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

// Wrap returns an error annotating err with a stack trace at the point Wrap is called,
// and the supplied message. If err is nil, Wrap returns nil.
func Wrap(err error, message string) error {
	return errors.Wrap(err, message)
}

// Cause returns the underlying cause of the error, if possible
func Cause(err error) error {
	return errors.Cause(err)
}

// New returns an error with the supplied message.
// New also records the stack trace at the point it was called.
func New(message string) error {
	return errors.New(message)
}

const (
	// ErrNoDataFound is returned when data not found for a request.
	ErrNoDataFound = "No data found for the request"
	// ErrGettingGeoLocation ...
	ErrGettingGeoLocation = "error happend while getting geo locations from redis"
	// ErrSavingGeoLocation ...
	ErrSavingGeoLocation = "error happend while adding geo location to redis"
)
