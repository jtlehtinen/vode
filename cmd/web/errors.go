package main

import (
	"net/http"
)

func (app *application) errorResponse(w http.ResponseWriter, status int, message any) {
	data := map[string]any{
		"error": message,
	}

	err := app.writeJSON(w, status, data, nil)
	if err != nil {
		app.logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	app.logger.Println(err)

	message := "server encountered an unexpected error and could not process the request"
	app.errorResponse(w, http.StatusInternalServerError, message)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	message := http.StatusText(status)
	app.errorResponse(w, status, message)
}

func (app *application) notFound(w http.ResponseWriter) {
	message := http.StatusText(http.StatusNotFound)
	app.errorResponse(w, http.StatusNotFound, message)
}
