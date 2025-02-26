FROM golang:1.24-alpine as build
RUN apk --no-cache add make
ARG OTEL_BINARY_NAME=simple-otel-collector
WORKDIR /app
COPY . .
RUN apk --no-cache add --update bash  # Install bash if not present
RUN make setup
RUN make build

FROM alpine:3.14
RUN apk --no-cache add curl coreutils
COPY --from=build /app/bin/${OTEL_BINARY_NAME} ./${OTEL_BINARY_NAME}
COPY --from=build /app/otelcol-config.yaml ./otelcol-config.yaml

EXPOSE 4317
EXPOSE 4318

CMD ["./simple-otel-collector", "--config=otelcol-config.yaml"]