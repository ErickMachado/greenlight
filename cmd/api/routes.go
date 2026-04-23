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
	mux.HandlerFunc(http.MethodPatch, "/v1/movies/:id", app.updateMovieHandler)
	mux.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.deleteMovieHandler)

	return app.recoverPanic(mux)
}
