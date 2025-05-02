package daemon

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

// Listen starts the HTTP server which handles connections.
// Cancelling the context will start a graceful shutdown of the server. The
// server continues to run until all remaining active connections have been
// served or the [Config.ShutdownTimeout] is reached.
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
			err := srv.Shutdown(shutdownCtx)
			if err != nil {
				return fmt.Errorf("error shutting down: %w", err)
			}
			return nil
		case err := <-ch:
			return err
		}
	}
}
