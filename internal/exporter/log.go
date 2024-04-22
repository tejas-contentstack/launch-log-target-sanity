package exporter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"go.opentelemetry.io/collector/client"
	"go.opentelemetry.io/collector/pdata/plog"
)

type LogEntry struct {
	Token string `json:"token"`
	Log   string `json:"log"`
}

type LogExporter struct {
	logs        []LogEntry
	mutex       sync.RWMutex
	clearTicker *time.Ticker
}

const (
	clearInterval       = 30 * time.Second
	authorizationHeader = "log-target-secret"
)

func NewLogExporter() *LogExporter {
	logExporter := &LogExporter{
		clearTicker: time.NewTicker(clearInterval),
	}
	go logExporter.clearLogsPeriodically()
	return logExporter
}

func (logExporter *LogExporter) pushLogs(ctx context.Context, logs plog.Logs) error {
	logExporter.mutex.Lock()
	defer logExporter.mutex.Unlock()

	tokens := client.FromContext(ctx).Metadata.Get(authorizationHeader)
	var token string
	if len(tokens) > 0 {
		token = tokens[0]
	} else {
		token = ""
	}

	firstLogBody := logs.ResourceLogs().At(0).ScopeLogs().At(0).LogRecords().At(0).Body().AsString()

	logEntry := LogEntry{
		Token: token,
		Log:   firstLogBody,
	}
	logExporter.logs = append(logExporter.logs, logEntry)

	return nil
}

func (logExporter *LogExporter) GetLogsHandler(responseWriter http.ResponseWriter, request *http.Request) {
	logExporter.mutex.RLock()
	defer logExporter.mutex.RUnlock()

	jsonLogs, jsonMarshalErr := json.Marshal(logExporter.logs)
	if jsonMarshalErr != nil {
		http.Error(responseWriter, fmt.Sprintf("failed to marshal logs to JSON: %v", jsonMarshalErr), http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Write(jsonLogs)
}

func (logExporter *LogExporter) clearLogsPeriodically() {
	for range logExporter.clearTicker.C {
		logExporter.clearLogs()
	}
}

func (logExporter *LogExporter) clearLogs() {
	logExporter.mutex.Lock()
	defer logExporter.mutex.Unlock()

	logExporter.logs = nil
}
