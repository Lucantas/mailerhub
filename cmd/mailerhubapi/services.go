package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func handleServices() {
	r.Route("/v1/services", func(r chi.Router) {

		r.Route("/address-list", func(r chi.Router) {
			r.Use(ServiceCtx)
			r.Post("/", newAddressList)
		})

		r.Route("/address-list/{serviceId}", func(r chi.Router) {
			r.Use(ServiceCtx)
			r.Get("/", getAddressInfo)
			r.Put("/", editAddress)
			r.Delete("/", deleteAddress)
			r.Post("/", newAddress)
		})

		/* r.Route("/{serviceId}", func(r chi.Router) {
			r.Use(ServiceCtx)
			r.Get("/", getService)
			r.Put("/", updateService)
			r.Delete("/", deleteService)
			r.Post("/", fireService)
		}) */
	})
}

func getAddresses(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Address"))
}

func getAddressInfo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Address"))
}

func editAddress(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Address"))
}

func deleteAddress(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Address"))
}

func newAddress(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Address"))
}

// if
func newAddressList(w http.ResponseWriter, r *http.Request) {
	log.Println("new address and new list")
}

// ServiceCtx is the middleware of the services context
// every processing of the request about services must be
func ServiceCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO : Handle services authentication
		serviceId, err := strconv.Atoi(chi.URLParam(r, "serviceId"))

		if err != nil {
			makeResponse(err.Error(), w, 422)
			return
		}

		log.Println(serviceId)
		//service, err := mailerdb.GetService(serviceId)

		ctx := context.WithValue(r.Context(), serviceId, "serviceId")
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
