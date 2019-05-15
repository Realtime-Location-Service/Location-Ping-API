package location

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
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

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Method("POST", "/", saveLocationHandler)
		r.Method("GET", "/", getLocationHandler)
	})

	return r
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; utf8")
	res := response.(*locationResponse)
	if res.Err != nil {
		w.WriteHeader(res.Err.StatusCode)
	}
	return json.NewEncoder(w).Encode(res)
}

func decodeSaveLocationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req locationRequest

	// TODO: need to handle invalid types
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
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
