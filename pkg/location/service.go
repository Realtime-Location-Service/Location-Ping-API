package location

import (
	"context"
	"net/http"

	"github.com/rls/ping-api/store/repo"
	"github.com/rls/ping-api/utils/errors"
)

// Service is the interface that provides recipes.
type Service interface {
	Save(context.Context, *locationRequest) (*locationResponse, error)
}

type service struct {
	repo repo.ILocation
}

// Save method saves geo locations
func (svc *service) Save(ctx context.Context, r *locationRequest) (*locationResponse, error) {
	if err := svc.repo.Save(r.Referrer, r.Locations); err != nil {
		return &locationResponse{nil, errors.NewErr(http.StatusBadRequest, errors.Cause(err).Error())}, err
	}
	return &locationResponse{"Successfully saved locations", nil}, nil
}

// NewService creates a location service with necessary dependencies.
func NewService(repo repo.ILocation) Service {
	return &service{repo}
}
