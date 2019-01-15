package main

import (
	"encoding/json"
	"mailer-service/internal/pkg/apihelper/jwt"
	"mailer-service/internal/pkg/mailercore/mailerdb"
	"net/http"
)

func authorized(token string) bool {

	if jwt.ValidateToken(token) {
		return true
	}

	return false

}

func authenticate(login string, pass string) string {
	dbUser := mailerdb.GetUserByLogin(login)

	var token string

	if dbUser.MatchPassword(pass) {
		// generate and return token
		token = jwt.GenerateToken(
			dbUser.Name,
			dbUser.Email,
			dbUser.ID,
		)
		return token
	}

	// handle wrong passwords
	return token
}

func (a api) authHandler(w http.ResponseWriter, r *http.Request) {

	var aB authBody

	if r.Body == nil {
		makeResponse("Data not found", json.NewEncoder(w), 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&aB)

	if err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		makeResponse("Unprocessable Entity", json.NewEncoder(w), 422)
		return
	}

	token := authenticate(aB.Login, aB.Password)

	if token == "" {
		makeResponse("Unprocessable Entity", json.NewEncoder(w), 422)
		return
	}

	makeResponse(token, json.NewEncoder(w), 200)
}

type authBody struct {
	Login    string
	Password string
}
