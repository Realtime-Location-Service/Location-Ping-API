package location

import (
	"context"
	"net/http"

	"github.com/asaskevich/govalidator"

	"github.com/rls/ping-api/store/model"
	"github.com/rls/ping-api/utils/errors"
)

type validationMiddleware struct {
	svc Service
}

// NewValidationMiddleware returns an instance of a location validation middleware.
func NewValidationMiddleware(svc Service) Service {
	defineCustomTags()
	return &validationMiddleware{svc}
}

func defineCustomTags() {
	govalidator.CustomTypeTagMap.Set("validate_locations", govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		locations := i.([]*model.Location)

		for _, loc := range locations {
			if ok, _ := govalidator.ValidateStruct(loc); !ok {
				return false
			}
		}
		return true
	}))
}

func (vmw validationMiddleware) Save(ctx context.Context, lr *locationRequest) (*locationResponse, error) {
	if _, err := govalidator.ValidateStruct(lr); err != nil {
		return &locationResponse{Data: nil, Err: errors.NewErr(http.StatusBadRequest, err.Error())}, nil
	}
	return vmw.svc.Save(ctx, lr)
}
