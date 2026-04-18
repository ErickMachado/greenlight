package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	mux := httprouter.New()

	mux.NotFound = http.HandlerFunc(app.notFoundResponse)
	mux.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	mux.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	mux.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	mux.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	return app.recoverPanic(mux)
}
