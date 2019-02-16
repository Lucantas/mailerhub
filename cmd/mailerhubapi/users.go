package main

import (
	"context"
	"encoding/json"
	"fmt"
	"mailer-service/internal/pkg/apihelper"
	"mailer-service/internal/pkg/mailercore/components"
	"mailer-service/internal/pkg/mailercore/mailerdb"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func handleUsers() {
	r.Route("/v1/users", func(r chi.Router) {

		r.Post("/", createUser)

		r.Route("/{userId}", func(r chi.Router) {
			r.Use(UserCtx)
			r.Use(crossCtx)
			r.Use(generalCtx)
			r.Get("/", getUser)
			r.Put("/", updateUser)
			r.Delete("/", deleteUser)
		})
	})
}

// UserCtx is the middleware of the users context
// every processing of the request about users must be
func UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		/* token := r.Header.Get("Mailer-Authentication")

		if !authorized(token) {
			http.Error(w, http.StatusText(403), 403)
			return
		} */

		userID, err := strconv.Atoi(chi.URLParam(r, "userId"))

		if err != nil {
			http.Error(w, http.StatusText(422), 422)
		}

		user, err := mailerdb.GetUser(userID)

		if err != nil {
			if mailerdb.IsErrUserNotFound(err) {
				makeResponse(err.Error(), w, 404)
				return
			}
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user components.User

	if r.Body == nil {
		makeResponse("Data not found", w, 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		makeResponse("Unprocessable Entity", w, 422)
		return
	}

	user.Password = apihelper.SecurePassword(user.Password)

	if user.Email == "" {
		makeResponse("Invalid email", w, 422)
		return
	}

	if user.Name == "" {
		makeResponse("Name is mandatory", w, 422)
	}

	err = mailerdb.SaveUser(user)

	if err != nil {
		makeResponse(
			fmt.Sprintf("Error while creating user: %s", err),
			w,
			403,
		)
		return
	}

	makeResponse("User successfully stored", w, 201)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get users"))
}

func getUser(w http.ResponseWriter, r *http.Request) {

	userID, err := strconv.Atoi(chi.URLParam(r, "userId"))

	if err != nil {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	u, err := mailerdb.GetUser(userID)

	if err != nil {
		if err != nil {
			if _, ok := err.(mailerdb.ErrUserNotFound); ok {
				makeResponse(fmt.Sprintf("%e", err), w, 404)
			}
		}
	}

	e := json.NewEncoder(w)

	e.Encode(u)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("udpate user by Id " + chi.URLParam(r, "userId")))
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete user by Id " + chi.URLParam(r, "userId")))
}
