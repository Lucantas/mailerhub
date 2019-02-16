package main

import (
	"encoding/json"
	"net/http"
)

//response is a basic struct cotaining a name and statuscode to an error
type response struct {
	Message    string `json:"message"`
	StatusCode int
}

func makeResponse(n string, w http.ResponseWriter, sCode int) {

	e := json.NewEncoder(w)

	msg := response{Message: n, StatusCode: sCode}

	w.WriteHeader(sCode)

	e.Encode(msg)
}
