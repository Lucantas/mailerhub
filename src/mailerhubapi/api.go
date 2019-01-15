package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

var r = chi.NewRouter()

type api struct {
}

func (a api) run() {
	r.Get("/authenticate", a.authHandler)
	a.handleUsers()

	//r.Get("/v1/users", a.getUsers)
	//r.Get("/v1/users/{userId}", a.getUserByID)
	//a.handleServices()
	http.ListenAndServe(":3000", r)
}

/* func (a api) handleServices() {

	r.Get("/v1/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome services"))
	})

}
*/