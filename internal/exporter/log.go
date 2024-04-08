package exporter

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/client"
	"go.opentelemetry.io/collector/pdata/plog"
)

type LogExporter struct {
}

func NewLogExporter() *LogExporter {
	return &LogExporter{}
}

func (logExporter *LogExporter) pushLogs(ctx context.Context, logs plog.Logs) error {
	fmt.Println("headers-->", client.FromContext(ctx).Metadata)
	// Create a JSON marshaler
	marshaler := &plog.JSONMarshaler{}
	// Marshal logs to JSON bytes
	jsonBytes, err := marshaler.MarshalLogs(logs)
	if err != nil {
		return fmt.Errorf("failed to marshal logs to JSON: %v", err)
	}
	// Print JSON logs
	fmt.Println(string(jsonBytes))
	return nil
}
