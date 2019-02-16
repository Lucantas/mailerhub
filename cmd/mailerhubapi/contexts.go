package main

import (
	"context"
	"mailer-service/internal/pkg/apihelper"
	"net/http"
)

func crossCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// TODO: create a method to deal with cross origin requisitions properly
		r.Header.Set("Access-Control-Allow-Origin", "*")

		ctx := context.WithValue(r.Context(), "headers", "HEADERS")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func generalCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// TODO: create a method to deal with cross origin requisitions properly
		apihelper.SetHeaders(w)

		if r.Method == "OPTIONS" {

			return
		}

		ctx := context.WithValue(r.Context(), "headers", "HEADERS")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
