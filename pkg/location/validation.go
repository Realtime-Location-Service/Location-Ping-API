package location

import (
	"context"
	"net/http"

	"github.com/asaskevich/govalidator"

	"github.com/rls/ping-api/utils/errors"
)

type validationMiddleware struct {
	svc Service
}

// NewValidationMiddleware returns an instance of a location validation middleware.
func NewValidationMiddleware(svc Service) Service {
	return &validationMiddleware{svc}
}

func (vmw validationMiddleware) Save(ctx context.Context, lr *locationRequest) (*locationResponse, error) {
	if _, err := govalidator.ValidateStruct(lr); err != nil {
		return &locationResponse{Data: nil, Err: errors.NewErr(http.StatusBadRequest, err.Error())}, nil
	}
	return vmw.svc.Save(ctx, lr)
}
