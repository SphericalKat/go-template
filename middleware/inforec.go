package middleware

import (
	"context"
	"sample/logger"

	ogenmw "github.com/ogen-go/ogen/middleware"
)

func InfoRec() ogenmw.Middleware {
	return func(req ogenmw.Request, next ogenmw.Next) (ogenmw.Response, error) {
		ctx := req.Context
		ctx = context.WithValue(ctx, logger.Host, req.Raw.Host)
		ctx = context.WithValue(ctx, logger.Method, req.Raw.Method)
		ctx = context.WithValue(ctx, logger.Method, req.Raw.RequestURI)
		ctx = context.WithValue(ctx, logger.OperationName, req.OperationName)
		ctx = context.WithValue(ctx, logger.UserAgent, req.Raw.UserAgent())

		req.SetContext(ctx)
		resp, err := next(req)
		return resp, err
	}
}
