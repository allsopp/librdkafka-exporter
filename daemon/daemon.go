package daemon

import (
	"context"
	"fmt"
	"io"
	"log"
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
// functionality, including the POST /metrics endpoint, which uses the Read()
// method to accept metric values in JSON format, and the GET /metrics
// endpoint, which uses the Handler() method to returns metric values in
// Prometheus exposition format. See [metrics.Metrics] for the reference
// implementation of this interface.
type Metrics interface {
	Read(io.Reader) error
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

// Listen starts the HTTP server which handles connections.
//
// Cancelling the context will start a graceful shutdown of the server. The
// server continues to run until all remaining active connections have been
// served or the Config.ShutdownTimeout is reached.
func (d Daemon) Listen(ctx context.Context, cfg Config) error {
	srv := &http.Server{
		Handler:      d.router,
		Addr:         cfg.ListenAddr,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	ch := make(chan error)

	go func() {
		log.Printf("listening on %s", cfg.ListenAddr)
		ch <- srv.ListenAndServe()
	}()

	for {
		select {
		case <-ctx.Done():
			log.Printf("shutting down gracefully...")
			shutdownCtx, cancel := context.WithTimeout(
				context.Background(),
				cfg.ShutdownTimeout,
			)
			defer cancel()
			return fmt.Errorf("shutdown: %w", srv.Shutdown(shutdownCtx))
		case err := <-ch:
			return err
		}
	}
}

func (d Daemon) metricsHandler(w http.ResponseWriter, r *http.Request) {
	d.metrics.Handler().ServeHTTP(w, r)
}

func (d Daemon) updateHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	err := d.metrics.Read(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
