package middleware

import (
	"sample/logger/tlog"
	"sample/utils"

	"github.com/go-errors/errors"
	ogenmw "github.com/ogen-go/ogen/middleware"
)

func PanicCapture() ogenmw.Middleware {
	return func(req ogenmw.Request, next ogenmw.Next) (resp ogenmw.Response, respErr error) {
		defer func() {
			if err := recover(); err != nil {
				err := errors.Wrap(err, 1)
				tlog.Error(req.Context, utils.GetStack(err), nil)

				// TODO: make the resp say something like "internal error"
				resp = ogenmw.Response{}
				respErr = err
			}
		}()

		resp, respErr = next(req)
		return
	}
}
