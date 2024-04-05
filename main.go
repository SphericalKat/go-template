package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	genapi "sample/api/gen"
	"sample/middleware"
	"sample/pkg/di"
	"sample/utils"
	"time"

	"github.com/fatih/color"
	"github.com/go-errors/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
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
	exporter, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint("jaeger:4318"), otlptracehttp.WithInsecure())
	if err != nil {
		handleErr(err)
	}

	resource, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String("gotemplate"),
			semconv.ServiceVersionKey.String("0.1.0"),
			attribute.String("environment", "local"),
		),
	)
	if err != nil {
		handleErr(err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource),
	)

	// ─── start http server ────────────────────────────────────────────────
	handler, err := di.InjectHandlers()
	if err != nil {
		handleErr(err)
	}

	var server *genapi.Server
	if server, err = genapi.NewServer(
		handler,
		genapi.WithTracerProvider(tp),
		genapi.WithMiddleware(middleware.InfoRec()),
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
		os.Exit(1)
	}
}

func handleErr(err error) {
	tperr := errors.Wrap(err, 0)
	fmt.Printf("%s", utils.PrintStack(tperr))
	panic(err)
}
