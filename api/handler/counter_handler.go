package handler

import (
	"context"
	"fmt"

	genapi "github.com/goalistdeveloper/go-template/api/gen"
	"github.com/goalistdeveloper/go-template/logger/tlog"
	"github.com/goalistdeveloper/go-template/pkg/domain/service"
)

type CounterHandler interface {
	GetCurrentCount(ctx context.Context) (*genapi.Counter, error)
	IncrementCount(ctx context.Context, req genapi.OptIncrement) (*genapi.Counter, error)
}

type counterHandlerImpl struct {
	counterService service.CounterService
}

// TODO: properly implement the two functions
func (h *counterHandlerImpl) GetCurrentCount(ctx context.Context) (*genapi.Counter, error) {
	tlog.Info(ctx, "Testing logs in trace", &map[string]any{
		"meaningOfLife": "42",
	})

	var a []string
	fmt.Println(a[3])

	return &genapi.Counter{
		Count: 0,
	}, nil
}

func (h *counterHandlerImpl) IncrementCount(ctx context.Context, req genapi.OptIncrement) (*genapi.Counter, error) {
	return &genapi.Counter{
		Count: 1,
	}, nil
}
