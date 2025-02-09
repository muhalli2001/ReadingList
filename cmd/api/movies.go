package main
import (
"fmt"
//"encoding/json"
"net/http"
"time" // new import
"github.com/muhalli2001/ReadingList/internal/data" // New import
)

// post method (it says that in routes.go but for clarity)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
	Title string `json:"title"`
	Year int32 `json:"year"`
	Runtime data.Runtime `json:"runtime"` // Make this field a data.Runtime type.
	Genres []string `json:"genres"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
	app.badRequestResponse(w, r, err)
	return
	}
	fmt.Fprintf(w, "%+v\n", input)
	}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
	// Use the new notFoundResponse() helper.
	app.notFoundResponse(w, r)
	return
	}
	movie := data.Movie{
	ID: id,
	CreatedAt: time.Now(),
	Title: "Casablanca",
	Runtime: 102,
	Genres: []string{"drama", "romance", "war"},
	Version: 1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
	// Use the new serverErrorResponse() helper.
	app.serverErrorResponse(w, r, err)
	}
	}
	