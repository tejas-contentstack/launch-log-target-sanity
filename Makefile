OTEL_BINARY_PATH=bin/simple-otel-collector
CONFIG_FILE=otelcol-config.yaml
HTTP_ENDPOINT=0.0.0.0:4318
GRPC_ENDPOINT=0.0.0.0:4317
HEALTH_CHECK_ENDPOINT= 0.0.0.0:13133

.PHONY: setup
setup: 
	./generate_config.sh $(HTTP_ENDPOINT) $(GRPC_ENDPOINT) $(HEALTH_CHECK_ENDPOINT) && go mod tidy

.PHONY: build
build: 
	go build -o ${OTEL_BINARY_PATH} .

.PHONY: run
run: 
	go run . --config ${CONFIG_FILE}
