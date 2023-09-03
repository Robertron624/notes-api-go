package handlers

import "net/http"

func (h handler) HealthCheck(w http.ResponseWriter, r *http.Request) {

	// Send an OK response
	w.WriteHeader(http.StatusOK)

	// Send a response indicating that the service is healthy
	w.Write([]byte("Service is healthy"))

}
