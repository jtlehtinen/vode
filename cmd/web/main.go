package main

import (
	"log"
	"net/http"
	"os"
)

type Poll struct {
	ID         int64
	Question   string
	Options    []string
	Votes      []int
	TotalVotes int
}

type application struct {
	polls  map[int64]*Poll
	logger *log.Logger
}

func main() {
	logger := log.New(os.Stdout, "", log.Ltime|log.Lshortfile)

	app := &application{
		polls:  make(map[int64]*Poll),
		logger: logger,
	}

	srv := &http.Server{
		Addr:    "localhost:5000",
		Handler: app.routes(),
	}

	logger.Printf("Starting server on %s\n", srv.Addr)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
