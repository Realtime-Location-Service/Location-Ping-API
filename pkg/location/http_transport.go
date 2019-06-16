package location

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/rls/ping-api/utils/consts"
	"github.com/rls/ping-api/utils/errors"
)

// MakeHandler returns a handler for the location service.
func MakeHandler(svc Service) http.Handler {
	saveLocationHandler := kithttp.NewServer(
		makeSaveLocationEndpoint(svc),
		decodeSaveLocationRequest,
		encodeResponse,
	)

	getLocationHandler := kithttp.NewServer(
		makeGetLocationEndpoint(svc),
		decodeGetLocationRequest,
		encodeResponse,
	)

	searchLocationHandler := kithttp.NewServer(
		makeSearchLocationEndpoint(svc),
		decodeSearchLocationRequest,
		encodeResponse,
	)

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Method("POST", "/", saveLocationHandler)
		r.Method("GET", "/", getLocationHandler)
		r.Method("GET", "/users", searchLocationHandler)
	})

	return r
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", consts.JSONContent)
	res := response.(*locationResponse)
	if res.Err != nil {
		w.WriteHeader(res.Err.StatusCode)
	}
	return json.NewEncoder(w).Encode(res)
}

func decodeSaveLocationRequest(_ context.Context, r *http.Request) (i interface{}, e error) {
	var req locationRequest

	defer func() {
		if r := recover(); r != nil {
			i, e = &locationResponse{nil, errors.NewErr(http.StatusBadRequest, errors.ErrInvalidJSON)}, nil
		}
	}()

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return &locationResponse{nil, errors.NewErr(http.StatusBadRequest, err.Error())}, nil
	}
	req.Referrer = r.Header.Get("RLS-Referrer")
	req.Locations.UserID = req.UserID

	return req, nil
}

func decodeGetLocationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req getLocationRequest
	req.UserIDs = []string{}
	for _, id := range strings.Split(r.URL.Query().Get("user_ids"), ",") {
		if id = strings.TrimSpace(id); id != "" {
			req.UserIDs = append(req.UserIDs, id)
		}
	}
	req.Referrer = r.Header.Get("RLS-Referrer")
	return req, nil
}

func decodeSearchLocationRequest(_ context.Context, r *http.Request) (i interface{}, e error) {
	var req searchLocationRequest

	defer func() {
		if r := recover(); r != nil {
			i, e = &locationResponse{nil, errors.NewErr(http.StatusBadRequest, errors.ErrParsingQueryString)}, nil
		}
	}()

	q := r.URL.Query()
	base := 64

	req.Radius, _ = strconv.ParseFloat(strings.TrimSpace(q.Get("radius")), base)
	req.Lat, _ = strconv.ParseFloat(strings.TrimSpace(q.Get("lat")), base)
	req.Lon, _ = strconv.ParseFloat(strings.TrimSpace(q.Get("lon")), base)
	req.Unit = strings.TrimSpace(q.Get("unit"))
	req.Referrer = r.Header.Get("RLS-Referrer")

	req.Limit, _ = strconv.Atoi(strings.TrimSpace(q.Get("limit")))
	if req.Limit <= 0 {
		req.Limit = consts.DefaultLocationsLimit
	}

	return req, nil
}
