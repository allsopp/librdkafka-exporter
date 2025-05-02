/*
Package daemon implements a daemon to import librdkafka metrics from JSON and
export librdkafka metrics in Prometheus exposition format over an HTTP API. An
example command to run the daemon is provided at the root of this distribution.
*/
package daemon

import (
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

// Config holds configuration for the daemon.
type Config struct {
	ListenAddr      string        // Address and TCP port to listen on, e.g. "127.0.0.1:8080"
	ReadTimeout     time.Duration // Amount of time allowed to read request
	WriteTimeout    time.Duration // Amount of time allowed to write response
	ShutdownTimeout time.Duration // Amount of time allowed for graceful shutdown
}

// Metrics is the interface required for the daemon to support metrics
// functionality, including the POST /metrics endpoint, which uses the
// ReadFrom() method to accept metric values in JSON format, and the GET
// /metrics endpoint, which uses the Handler() method to returns metric values
// in Prometheus exposition format.
//
// See [github.com/allsopp/librdkafka-exporter/metrics.Metrics] for the
// reference implementation of this interface.
type Metrics interface {
	ReadFrom(io.Reader) error
	Handler() http.Handler
}

// Daemon holds internal state and contains no exported fields.
type Daemon struct {
	router  chi.Router
	metrics Metrics
}

// New returns an instance of the daemon.
func New(m Metrics) Daemon {
	daemon := Daemon{
		router:  chi.NewRouter(),
		metrics: m,
	}

	daemon.router.Use(middleware.StripSlashes)
	daemon.router.Use(middleware.Heartbeat("/health"))
	act := middleware.AllowContentType("application/json")
	daemon.router.With(act).Post("/metrics", daemon.updateHandler)
	daemon.router.Get("/metrics", daemon.metricsHandler)
	return daemon
}
