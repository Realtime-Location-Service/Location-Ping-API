package location

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/rls/ping-api/store/model"
	"github.com/rls/ping-api/utils/errors"
)

type locationRequest struct {
	UserID    string          `json:"user_id" valid:"required"`
	Referrer  string          `json:"referrer" valid:"required"`
	Locations *model.Location `json:"locations" valid:"required"`
}
type locationResponse struct {
	Data interface{} `json:"data,omitempty"`
	Err  *errors.Err `json:"err,omitempty"`
}

func makeSaveLocationEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(locationRequest)
		return svc.Save(ctx, &req)
	}
}
