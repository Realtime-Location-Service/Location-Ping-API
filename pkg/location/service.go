package location

import (
	"context"
	"net/http"

	"github.com/rls/ping-api/store/model"
	"github.com/rls/ping-api/store/repo"
	"github.com/rls/ping-api/utils/errors"
)

// Service is the interface that provides recipes.
type Service interface {
	Save(context.Context, *locationRequest) (*locationResponse, error)
	Get(context.Context, *getLocationRequest) (*locationResponse, error)
	Search(context.Context, *searchLocationRequest) (*locationResponse, error)
}

type service struct {
	repo repo.ILocation
}

// Save method saves geo locations
func (svc *service) Save(ctx context.Context, r *locationRequest) (*locationResponse, error) {
	if err := svc.repo.Save(r.Referrer, r.Locations); err != nil {
		return &locationResponse{nil, errors.NewErr(http.StatusBadRequest, err.Error())}, nil
	}
	return &locationResponse{"Successfully saved locations", nil}, nil
}

// Get method returns users geo locations
func (svc *service) Get(ctx context.Context, r *getLocationRequest) (*locationResponse, error) {
	locations, err := svc.repo.Get(r.Referrer, r.UserIDs)
	if err != nil {
		return &locationResponse{nil, errors.NewErr(http.StatusBadRequest, err.Error())}, nil
	}
	return &locationResponse{locations, nil}, nil
}

// Search returns users location within radius
func (svc *service) Search(ctx context.Context, r *searchLocationRequest) (*locationResponse, error) {
	locations, err := svc.repo.Search(r.Referrer,
		&model.Radius{Lat: r.Lat, Lon: r.Lon, Val: r.Radius, Unit: r.Unit, Limit: r.Limit})

	if err != nil {
		return &locationResponse{nil, errors.NewErr(http.StatusBadRequest, err.Error())}, nil
	}
	return &locationResponse{locations, nil}, nil
}

// NewService creates a location service with necessary dependencies.
func NewService(repo repo.ILocation) Service {
	return &service{repo}
}
