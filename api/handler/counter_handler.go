package handler

import (
	"context"
	genapi "sample/api/gen"
	"sample/pkg/domain/service"
)

type CounterHandler interface {
	GetCurrentCount(ctx context.Context) (*genapi.CounterResponse, error)
	IncrementCount(ctx context.Context, req genapi.OptIncrementRequest) (*genapi.CounterResponse, error)
}

type counterHandlerImpl struct {
	counterService service.CounterService
}

// TODO: properly implement the two functions
func (h *counterHandlerImpl) GetCurrentCount(ctx context.Context) (*genapi.CounterResponse, error) {
	return &genapi.CounterResponse{
		Count: 0,
	}, nil
}

func (h *counterHandlerImpl) IncrementCount(ctx context.Context, req genapi.OptIncrementRequest) (*genapi.CounterResponse, error) {
	return &genapi.CounterResponse{
		Count: 1,
	}, nil
}
