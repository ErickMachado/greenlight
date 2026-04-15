package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	mux := httprouter.New()

	mux.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	mux.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	mux.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	return mux
}
