package main

import (
	"context"
	"encoding/json"
	"fmt"
	"mailer-service/internal/pkg/mailercore/components"
	"mailer-service/internal/pkg/mailercore/mailerdb"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (a api) handleUsers() {
	r.Route("/v1/users", func(r chi.Router) {

		r.Post("/", a.createUser)

		r.Route("/{userId}", func(r chi.Router) {
			r.Use(UserCtx)
			r.Get("/", a.getUser)
			r.Put("/", a.updateUser)
			r.Delete("/", a.deleteUser)
		})
	})
}

// UserCtx is the middleware of the users context
// every processing of the request about users must be
func UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Mailer-Authentication")

		if !authorized(token) {
			http.Error(w, http.StatusText(403), 403)
			return
		}

		userID, err := strconv.Atoi(chi.URLParam(r, "userId"))

		if err != nil {
			http.Error(w, http.StatusText(422), 422)
		}

		user := mailerdb.GetUser(userID)

		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a api) createUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create user"))

	var user components.User

	if r.Body == nil {
		makeResponse("Data not found", json.NewEncoder(w), 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		makeResponse("Unprocessable Entity", json.NewEncoder(w), 422)
		return
	}

	mailerdb.SaveUser(user)
	makeResponse("User successfully stored", json.NewEncoder(w), 202)
}

func (a api) getUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get users"))
}

func (a api) getUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "userId"))
	if err != nil {
		http.Error(w, http.StatusText(422), 422)
	}
	u := mailerdb.GetUser(userID)
	fmt.Println("DEBUG> ", u.Name)
	w.Write([]byte("get user by Id " + u.Name))
}

func (a api) updateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("udpate user by Id " + chi.URLParam(r, "userId")))
}

func (a api) deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete user by Id " + chi.URLParam(r, "userId")))
}
