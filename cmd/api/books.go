package main

import (
	//"fmt"
	"net/http"
	//"net/url"
	"log"
	"io"
	"github.com/muhalli2001/ReadingList/internal/data" // New import data
	"github.com/muhalli2001/ReadingList/internal/validator" // validator
)

func (app *application) searchBookHandler(w http.ResponseWriter, r *http.Request) {


	// Eventually move the URL builder into its own util. Will be useful for fetching the book image covers in the future.
	// start by reading the user's query
	var input struct{
		Query string `json:"query"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
	app.badRequestResponse(w, r, err)
	return
	}

	search:= &data.Search{
		Query: input.Query,
	}

	// Initialize a new Validator.
	v := validator.New()

	if data.ValidateSearch(v, search); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
		}


	EncodedQuery := app.URLencoder(input.Query)
	// QueryURL := fmt.Sprintf("%s?q=%s%s", 
	// data.BooksearchURL, EncodedQuery ,data.Fields)

	QueryURL := data.BooksearchURL+"?q="+EncodedQuery+data.Fields


	// Fetch book data from Open Library API
	res, err := http.Get(QueryURL)
	if err != nil {
		log.Println("Error fetching data:", err)
		app.serverErrorResponse(w, r, err)
		return
	}
	defer res.Body.Close() // Ensure the body is closed

	// Read response body
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		app.serverErrorResponse(w, r, err)
		return
	}

	// Encapsulate response in envelope
	env := envelope{
		"encapsulatedData": string(responseBody), // Convert []byte to string
	}

	// Send JSON response
	err = app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
