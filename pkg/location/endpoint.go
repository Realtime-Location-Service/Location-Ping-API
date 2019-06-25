package location

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/rls/ping-api/store/model"
	"github.com/rls/ping-api/utils/errors"
)

type locationRequest struct {
	Referrer  string            `json:"referrer" valid:"required"`
	Locations []*model.Location `json:"locations" valid:"required"`
}

type getLocationRequest struct {
	UserIDs  []string `json:"user_ids" valid:"required"`
	Referrer string   `json:"referrer" valid:"required"`
}

type searchLocationRequest struct {
	Lat      float64 `json:"lat" valid:"required,latitude"`
	Lon      float64 `json:"lon" valid:"required,longitude"`
	Radius   float64 `json:"radius" valid:"required"`
	Unit     string  `json:"unit" valid:"radius_unit_tag~unit: invalid radius unit, required"`
	Referrer string  `json:"referrer" valid:"required"`
	Limit    int     `json:"limit" valid:"-"`
}

type locationResponse struct {
	Data interface{} `json:"data,omitempty"`
	Err  *errors.Err `json:"err,omitempty"`
}

func makeSaveLocationEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(locationRequest)
		if !ok {
			return request, nil
		}
		return svc.Save(ctx, &req)
	}
}

func makeGetLocationEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getLocationRequest)
		return svc.Get(ctx, &req)
	}
}

func makeSearchLocationEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(searchLocationRequest)
		if !ok {
			return request, nil
		}
		return svc.Search(ctx, &req)
	}
}
