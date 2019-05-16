package location

import (
	"context"
	"net/http"

	"github.com/asaskevich/govalidator"

	"github.com/rls/ping-api/utils/consts"
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
	govalidator.CustomTypeTagMap.Set("radius_unit_tag", govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		unit := i.(string)
		_, ok := consts.ValidRadiusUnits[unit]
		return ok
	}))
}

func (vmw validationMiddleware) Save(ctx context.Context, lr *locationRequest) (*locationResponse, error) {
	if _, err := govalidator.ValidateStruct(lr); err != nil {
		return &locationResponse{Data: nil, Err: errors.NewErr(http.StatusBadRequest, err.Error())}, nil
	}
	return vmw.svc.Save(ctx, lr)
}

func (vmw validationMiddleware) Get(ctx context.Context, lr *getLocationRequest) (*locationResponse, error) {
	if _, err := govalidator.ValidateStruct(lr); err != nil {
		return &locationResponse{Data: nil, Err: errors.NewErr(http.StatusBadRequest, err.Error())}, nil
	}
	return vmw.svc.Get(ctx, lr)
}

func (vmw validationMiddleware) Search(ctx context.Context, lr *searchLocationRequest) (*locationResponse, error) {
	if _, err := govalidator.ValidateStruct(lr); err != nil {
		return &locationResponse{Data: nil, Err: errors.NewErr(http.StatusBadRequest, err.Error())}, nil
	}
	return vmw.svc.Search(ctx, lr)
}
