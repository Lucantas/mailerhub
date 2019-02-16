package main

import (
	"encoding/json"
	"log"
	"mailer-service/internal/pkg/apihelper"
	"mailer-service/internal/pkg/jwt"
	"mailer-service/internal/pkg/mailercore/mailerdb"
	"net/http"
)

func authorized(token string) bool {

	if ok, _ := jwt.ValidateToken(token); ok {
		return true
	}

	return false

}

func authenticate(login string, pass string) string {
	dbUser, err := mailerdb.GetUserByLogin(login)

	if err != nil {
		log.Println(err)
	}

	var token string

	if apihelper.PasswordMatch(dbUser.Password, pass) {
		// generate and return token
		token, err = jwt.GenerateToken(
			dbUser.Name,
			dbUser.Email,
			dbUser.ID,
		)

		if err != nil {
			return ""
		}

		return token
	}

	// handle wrong passwords
	return token
}

func (a api) authHandler(w http.ResponseWriter, r *http.Request) {

	var aB authBody

	if r.Body == nil {
		makeResponse("Data not found", w, 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&aB)

	if err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		makeResponse("Unprocessable Entity", w, 422)
		return
	}

	token := authenticate(aB.Login, aB.Password)

	if token == "" {
		makeResponse("Unprocessable Entity", w, 422)
		return
	}

	makeResponse(token, w, 200)
}

type authBody struct {
	Login    string
	Password string
}
