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
	var err error

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
