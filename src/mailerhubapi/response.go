package main

import (
	"encoding/json"
)

//response is a basic struct cotaining a name and statuscode to an error
type response struct {
	Message    string
	StatusCode int
}

func makeResponse(n string, e *json.Encoder, sCode int) {

	msg := response{Message: n, StatusCode: sCode}

	e.Encode(msg)

}
