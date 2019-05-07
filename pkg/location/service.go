package location

import (
	"context"
	"net/http"

	"github.com/rls/ping-api/store/model"
	"github.com/rls/ping-api/store/repo"
	errors "github.com/rls/ping-api/utils/error"
)

// Service is the interface that provides recipes.
type Service interface {
	Save(context.Context, *locationRequest) (*locationResponse, error)
}

type service struct {
	repo repo.ILocation
}

// Save method saves latlong into redis
// TODO: need to implement
func (svc *service) Save(ctx context.Context, r *locationRequest) (*locationResponse, error) {
	if err := svc.repo.Save(r.UserID, &model.Location{}); err != nil {
		return &locationResponse{nil, errors.NewErr(http.StatusBadRequest, "Error")}, nil
	}
	return &locationResponse{"Done", nil}, nil
}

// NewService creates a location service with necessary dependencies.
func NewService(repo repo.ILocation) Service {
	return &service{repo}
}
