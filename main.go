package main

import (
	"github.com/rs/zerolog/log"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/otelcol"
)

func main() {
	info := component.BuildInfo{
		Command:     "simple-otel-collector",
		Description: "Local OpenTelemetry Collector binary",
		Version:     "1.0.0",
	}
	factories, err := components()
	if err != nil {
		log.Fatal().Err(err)
	}
	settings := otelcol.CollectorSettings{
		BuildInfo: info,
		Factories: func() (otelcol.Factories, error) {
			return factories, nil
		},
	}

	if err := runInteractive(settings); err != nil {
		log.Fatal().Err(err)
	}
}

func runInteractive(params otelcol.CollectorSettings) error {
	cmd := otelcol.NewCommand(params)
	if err := cmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("collector server run finished with error:")
	}

	return nil
}
