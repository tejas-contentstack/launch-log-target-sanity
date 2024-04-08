package exporter

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/pdata/plog"
)

type LogExporter struct {
}

func NewLogExporter() *LogExporter {
	return &LogExporter{}
}

func (logExporter *LogExporter) pushLogs(ctx context.Context, logs plog.Logs) error {
	fmt.Println("test")
	return nil
}
