package location

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/rls/ping-api/store/model"
	errors "github.com/rls/ping-api/utils/error"
)

type locationRequest struct {
	model.Location
}

type locationResponse struct {
	Data interface{} `json:"data,omitempty"`
	Err  *errors.Err `json:"err,omitempty"`
}

func makeSaveLocationEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(locationRequest)
		v, err := svc.Save(ctx, &req)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}
