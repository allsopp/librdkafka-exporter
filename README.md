# librdkafka-exporter

[![Build](https://github.com/allsopp/librdkafka-exporter/actions/workflows/build.yml/badge.svg)](https://github.com/allsopp/librdkafka-exporter/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/allsopp/librdkafka-exporter.svg)](https://pkg.go.dev/github.com/allsopp/librdkafka-exporter)

Prometheus exporter for [librdkafka](https://github.com/confluentinc/librdkafka).

This repository includes:

 - A Go library to export librdkafka stats in Prometheus exposition format.
 - A daemon that exposes the same functionality over an HTTP API.

## Library

To get the library, run:

    go get github.com/allsopp/librdkafka-exporter

The reference documentation for the latest release is available [here](https://pkg.go.dev/github.com/allsopp/librdkafka-exporter/metrics).

## Daemon

First, clone this repository:

    git clone https://github.com/allsopp/librdkafka-exporter

Then run the daemon:

    cd librdkafka-exporter
    go run cmd/main.go

Or build an executable and run that:

    cd librdkafka-exporter
    go build -o main cmd/main.go
    ./main

Alternatively, an example `Dockerfile` for building a Docker container image to run the
daemon is included with this distribution.

### Endpoints

The daemon exposes the following HTTP endpoints:

#### POST /metrics

Making a `POST /metrics` request will update the metric values.

* The request body must be the JSON data returned from the
  [`stats_cb` callback of librdkafka](https://github.com/confluentinc/librdkafka/blob/master/STATISTICS.md).

* The content type must be `application/json`, otherwise a `415 Unsupported Media Type`
  status will be returned and the metrics will not be updated.

#### GET /metrics

Making a `GET /metrics` request will return the metric values in
[Prometheus exposition format](https://github.com/prometheus/docs/blob/main/content/docs/instrumenting/exposition_formats.md).

#### GET /health

Making a `GET /health` request will always return a `200 OK` if the daemon is running.
This is expected to be used for diagnostics and/or health probes.

### Configuration

The daemon can be configured with the following environment variables:

| Environment variable | Description |
| --- | --- |
| `LISTEN_ADDR`        | The address and TCP port to listen on (defaults to `:8080`)     |
| `READ_TIMEOUT`       | Amount of time allowed to read request (defaults to `30s`)      |
| `WRITE_TIMEOUT`      | Amount of time allowed to write response (defaults to `30s`)    |
| `SHUTDOWN_TIMEOUT`   | Amount of time allowed for graceful shutdown (defaults to `5m`) |

## See also

- [librdkafka statistics](https://docs.confluent.io/platform/current/clients/librdkafka/html/md_STATISTICS.html)
- [KIP-714: Client metrics and observability](https://cwiki.apache.org/confluence/display/KAFKA/KIP-714%3A+Client+metrics+and+observability)
- [librdkafka-prometheus-exporter](https://github.com/mcolomerc/librdkafka-prometheus-exporter), a similar project
