package daemon

import (
	"fmt"
	"net/http"
)

func (d Daemon) metricsHandler(w http.ResponseWriter, r *http.Request) {
	d.metrics.Handler().ServeHTTP(w, r)
}

func (d Daemon) updateHandler(w http.ResponseWriter, r *http.Request) {
	err := d.metrics.ReadFrom(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
