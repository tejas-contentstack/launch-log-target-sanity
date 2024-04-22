package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
	loggingexporter "github.com/venky0195/simple-otel-collector/internal/exporter"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/otelcol"
)

func main() {
	info := component.BuildInfo{
		Command:     "simple-otel-collector",
		Description: "Local OpenTelemetry Collector binary",
		Version:     "1.0.0",
	}

	logExporter := loggingexporter.NewLogExporter()

	factories, err := components(logExporter)
	if err != nil {
		log.Fatal().Err(err)
	}

	settings := otelcol.CollectorSettings{
		BuildInfo: info,
		Factories: func() (otelcol.Factories, error) {
			return factories, nil
		},
	}

	go func() {
		if runHTTPServerErr := runHTTPServer(logExporter); err != nil {
			log.Error().Err(runHTTPServerErr).Msg("failed to start HTTP server")
		}
	}()

	if runInteractiveErr := runInteractive(settings); err != nil {
		log.Fatal().Err(runInteractiveErr)
	}
}

func runInteractive(params otelcol.CollectorSettings) error {
	cmd := otelcol.NewCommand(params)
	if err := cmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("collector server run finished with error:")
	}

	return nil
}

func runHTTPServer(logExporter *loggingexporter.LogExporter) error {
	http.HandleFunc("/", logExporter.GetLogsHandler)
	httpServer := &http.Server{
		Addr: ":8080",
	}

	go func() {
		log.Info().Str("address", httpServer.Addr).Msg("HTTP server started and listening")
		if httpServerErr := httpServer.ListenAndServe(); httpServerErr != nil && httpServerErr != http.ErrServerClosed {
			log.Error().Err(httpServerErr).Msg("HTTP server failed")
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Info().Msg("Shutting down HTTP server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if shutDownErr := httpServer.Shutdown(ctx); shutDownErr != nil {
		log.Error().Err(shutDownErr).Msg("HTTP server shutdown error")
	}

	log.Info().Msg("HTTP server shutdown completed.")
	return nil
}
