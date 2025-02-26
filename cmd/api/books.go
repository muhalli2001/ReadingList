package main

import (
	//"fmt"
	"net/http"
	//"net/url"
	"log"
	"io"
	"github.com/muhalli2001/ReadingList/internal/data" // New import data
	"github.com/muhalli2001/ReadingList/internal/validator" // validator
	"time"
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


	// Fetch book data from Open Library API, using a client to enforce a timeout in case the API is unavailable.
	client := &http.Client{Timeout: 5 * time.Second} // 5-second timeout
	res, err := client.Get(QueryURL)
	if err != nil {
		log.Println("Error fetching data:", err)
		app.serverErrorResponse(w, r, err)
		return
	}
	defer res.Body.Close() // Ensure the body is closed

	// see if this can be encapsulated into something we have so far.
	// for now, it should display any non 200 or 500 errors.
	// will need to figure out how to test this.
	if res.StatusCode != http.StatusOK {
		log.Printf("API request failed: %s", res.Status)
		app.errorResponse(w, r, res.StatusCode, "Failed to fetch book data")
		return
	}

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
