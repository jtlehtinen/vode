package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func createPollID() int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63()
}

func filter(values []string, predicate func(string) bool) []string {
	r := make([]string, 0)
	for _, value := range values {
		if predicate(value) {
			r = append(r, value)
		}
	}
	return r
}

func (app *application) handleCreatePoll(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Title   string   `json:"title"`
		Multi   bool     `json:"multi"`
		Options []string `json:"options"`
	}

	type response struct {
		Title   string   `json:"title"`
		Multi   bool     `json:"multi"`
		Options []string `json:"options"`
	}

	var req request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	req.Options = filter(req.Options, func(option string) bool { return len(option) > 0 })

	if req.Title == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	if len(req.Options) < 2 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	poll := &Poll{
		ID:       createPollID(),
		Question: req.Title,
		Options:  req.Options,
		Votes:    make([]int, len(req.Options)),
	}

	app.polls[poll.ID] = poll

	headers := http.Header{}
	headers.Set("Location", fmt.Sprintf("/v1/polls/%d", poll.ID))

	err = app.writeJSON(w, http.StatusCreated, poll, headers)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) handleVote(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	if id == "" {
		app.notFound(w)
		return
	}

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	p, ok := app.polls[idInt]
	if !ok {
		app.notFound(w)
		return
	}

	type request struct {
		Option string `json:"option"`
	}

	var req request
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	if req.Option == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	idx := -1
	for i, option := range p.Options {
		if req.Option == option {
			idx = i
			break
		}
	}

	if idx == -1 {
		app.notFound(w)
		return
	}

	p.Votes[idx]++

	// @TODO: send current result
}

func (app *application) handleGetPollResults(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	if id == "" {
		app.notFound(w)
		return
	}
}
