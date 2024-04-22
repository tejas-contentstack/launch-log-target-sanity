#!/bin/bash

if [ "$#" -ne 3 ]; then
    echo "Usage: $0 <HTTP_ENDPOINT> <GRPC_ENDPOINT> <HEALTH_CHECK_ENDPOINT>"
    exit 1
fi

HTTP_ENDPOINT=$1
GRPC_ENDPOINT=$2
HEALTH_CHECK_ENDPOINT=$3

cat > otelcol-config.yaml <<EOF
receivers:
  otlp:
    protocols:
      http:
        endpoint: $HTTP_ENDPOINT
        include_metadata: true
      grpc:
        endpoint: $GRPC_ENDPOINT
        include_metadata: true

exporters:
  exporter:

extensions:
  health_check:
    endpoint: $HEALTH_CHECK_ENDPOINT

service:
  telemetry:
    logs:
      output_paths: ["stdout"]
      error_output_paths: ["stderr"]
  extensions: [health_check]
  pipelines:
    logs:
      receivers: [otlp]
      processors: []
      exporters: [exporter]
EOF
