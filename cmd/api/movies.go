package main

import (
	"fmt"
	"net/http"
)

// createMovieHandler create a new movie in the database.
// - POST /v1/movies
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

// showMovieHandler shows the details of a movie.
// - GET /v1/movies/:id
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show the details of movie %d\n", id)
}
