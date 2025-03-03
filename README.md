# simple-otel-collector

This collector is built in Go and utilizes the OpenTelemetry framework for telemetry purposes.

## Setup

### Prerequisites

- Go installed version 1.24 >=

### Setup

Before running the service, set up the configuration file for the OpenTelemetry collector:

```bash
make setup
```

### Running the Service

To run the service, execute the following command:

```bash
make run
```

This will start the service using the configuration specified in 'otelcol-config.yaml'.
The HTTP server will be available at the endpoint `localhost:4318`
The GRPC server will be available at the endpoint `localhost:4317`
The healthcheck of this service will be available at the endpoint `localhost:13133`
The custom HTTP server will be available at the endpoint `localhost:8080`

### Configuration Files

- `otelcol-config.yaml`: This file is used to configure the Open Telemetry Collector itself.

### Custom HTTP Server (Port 8080)

The custom HTTP server running on port 8080 saves logs in memory for 30 seconds and returns the logs in JSON format.

To retrieve saved logs, send a GET request to `localhost:8080`.
The logs will be returned in JSON format.

#### Log Structure

Each log entry consists of two fields:

```javascript
{ 
    "token" : "The authorization token associated with the log entry",
    "log": "The log message itself." 
}
```
