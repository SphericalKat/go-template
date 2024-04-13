package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	genapi "github.com/goalistdeveloper/go-template/api/gen"
	"github.com/goalistdeveloper/go-template/middleware"
	"github.com/goalistdeveloper/go-template/pkg/di"
	"github.com/goalistdeveloper/go-template/utils"

	"github.com/fatih/color"
	"github.com/go-errors/errors"
)

func main() {
	ctx := context.Background()

	// ─── ent auto-migrate ────────────────────────────────────────────────
	client, err := di.InjectDb()
	if err != nil {
		handleErr(err)
	}

	if err := client.Schema.Create(ctx); err != nil {
		handleErr(err)
	}

	// ─── Configure Tracer Provider ────────────────────────────────────────
	tp, err := di.InjectTracerProvider()
	if err != nil {
		handleErr(err)
	}

	// ─── start http server ────────────────────────────────────────────────
	handlers, err := di.InjectHandlers()
	if err != nil {
		handleErr(err)
	}

	var server *genapi.Server
	if server, err = genapi.NewServer(
		handlers,

		// tracer
		genapi.WithTracerProvider(tp),

		// middleware
		genapi.WithMiddleware(
			middleware.PanicCapture(),
			middleware.InfoRec(),
		),
	); err != nil {
		handleErr(err)
	}

	httpServer := http.Server{
		ReadHeaderTimeout: time.Second,
		Addr:              ":9000",
		Handler:           server,
	}

	go func() {
		fmt.Printf("%s running server on %s\n", color.YellowString("[info]"), "localhost:9000")
		err := httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			handleErr(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		slog.Error("shutdown err", "err", err)
		//nolint:gocritic
		os.Exit(1)
	}
}

func handleErr(err error) {
	tperr := errors.Wrap(err, 0)
	fmt.Printf("%s", utils.GetStack(tperr, true))
	panic(err)
}
