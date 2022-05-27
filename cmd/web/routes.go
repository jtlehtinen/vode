package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	router := pat.New()

	router.Post("/v1/polls", http.HandlerFunc(app.handleCreatePoll))
	router.Post("/v1/polls/:id", http.HandlerFunc(app.handleVote))
	router.Get("/v1/polls/:id", http.HandlerFunc(app.handleGetPollResults))

	return router
}
