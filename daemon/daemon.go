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

type Config struct {
	ListenAddr      string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

type Metrics interface {
	Read(io.Reader) error
	Handler() http.Handler
}

type Daemon struct {
	router  chi.Router
	metrics Metrics
}

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
			shutdownCtx, _ := context.WithTimeout(
				context.Background(),
				cfg.ShutdownTimeout,
			)
			return srv.Shutdown(shutdownCtx)
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
