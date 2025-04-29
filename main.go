/*
Run a daemon to import and export librdkafka metrics over an HTTP API.

# Configuration

The daemon can be configured by setting environment variables, which are
documented in the following subsections.

## Listen address

The daemon listens on the configured LISTEN_ADDR, which should be specified in
the format "host:port", where port is a TCP port number. This defaults to ":8080"

## Timeouts

The daemon will use the values of READ_TIMEOUT, WRITE_TIMEOUT, and
SHUTDOWN_TIMEOUT as the time allowed for server reads, server writes, and
graceful shutdown, respectively. These values should be provided in a format
compatible with [time.ParseDuration], for example "500ms", "30s", or "1m".

The default values for the timeouts are:

  - READ_TIMEOUT defaults to `30s`
  - WRITE_TIMEOUT defaults to `30s`
  - SHUTDOWN_TIMEOUT defaults to `10s`
*/
package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/allsopp/librdkafka-exporter/daemon"
	"github.com/allsopp/librdkafka-exporter/metrics"
	"github.com/sethvargo/go-envconfig"
)

type config struct {
	ListenAddr      string        `env:"LISTEN_ADDR, default=:8080"`
	ReadTimeout     time.Duration `env:"READ_TIMEOUT, default=30s"`
	WriteTimeout    time.Duration `env:"WRITE_TIMEOUT, default=30s"`
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT, default=10s"`
}

func main() {
	m, err := metrics.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	var cfg config
	err = envconfig.Process(ctx, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	err = daemon.New(m).Listen(
		ctx,
		daemon.Config{
			ListenAddr:      cfg.ListenAddr,
			ReadTimeout:     cfg.ReadTimeout,
			WriteTimeout:    cfg.WriteTimeout,
			ShutdownTimeout: cfg.ShutdownTimeout,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
