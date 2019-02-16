package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

func handleServices() {
	r.Route("/v1/services", func(r chi.Router) {

		r.Route("/{serviceId}", func(r chi.Router) {
			r.Use(ServiceCtx)
			r.Get("/", getService)
			r.Put("/", updateService)
			r.Delete("/", deleteService)
			r.Post("/", fireService)
		})
	})
}

// ServiceCtx is the middleware of the services context
// every processing of the request about services must be
func ServiceCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO : Handle services authentication
		ctx := context.WithValue(r.Context(), "service", "")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getService(w http.ResponseWriter, r *http.Request) {
	// TODO: get service info from the base to display it to user
}

func updateService(w http.ResponseWriter, r *http.Request) {
	// TODO: update service info
}

func deleteService(w http.ResponseWriter, r *http.Request) {
	// TODO: delete service from base
}

func fireService(w http.ResponseWriter, r *http.Request) {
	// TODO: Send the email body to the address of the service
}
