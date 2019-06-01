package location

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	svc    Service
}

// NewLoggingMiddleware returns a new instance of a location logging middleware.
func NewLoggingMiddleware(logger log.Logger, svc Service) Service {
	return &loggingMiddleware{logger, svc}
}

func (lmw loggingMiddleware) Save(ctx context.Context, r *locationRequest) (output *locationResponse, err error) {

	defer func(begin time.Time) {
		req, _ := json.Marshal(r)
		resp, _ := json.Marshal(output.Data)
		err, _ := json.Marshal(output.Err)

		lmw.logger.Log(
			"method", "Save",
			"input", req,
			"output", resp,
			"error", err,
			"TimeTaken", time.Since(begin),
		)
	}(time.Now())

	output, err = lmw.svc.Save(ctx, r)
	return
}

func (lmw loggingMiddleware) Get(ctx context.Context, r *getLocationRequest) (output *locationResponse, err error) {

	defer func(begin time.Time) {
		req, _ := json.Marshal(r)
		resp, _ := json.Marshal(output.Data)
		err, _ := json.Marshal(output.Err)

		lmw.logger.Log(
			"method", "Get",
			"input", req,
			"output", resp,
			"error", err,
			"TimeTaken", time.Since(begin),
		)
	}(time.Now())

	output, err = lmw.svc.Get(ctx, r)
	return
}

func (lmw loggingMiddleware) Search(ctx context.Context, r *searchLocationRequest) (output *locationResponse, err error) {

	defer func(begin time.Time) {
		req, _ := json.Marshal(r)
		resp, _ := json.Marshal(output.Data)
		err, _ := json.Marshal(output.Err)

		lmw.logger.Log(
			"method", "Search",
			"input", req,
			"output", resp,
			"error", err,
			"TimeTaken", time.Since(begin),
		)
	}(time.Now())

	output, err = lmw.svc.Search(ctx, r)
	return
}
