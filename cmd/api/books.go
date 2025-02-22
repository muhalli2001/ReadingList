package main

import (
	//"fmt"
	"net/http"
	"log"
	"io"
)

func (app *application) searchBookHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch book data from Open Library API
	res, err := http.Get("https://openlibrary.org/search.json?q=the+lord+of+the+rings")
	if err != nil {
		log.Println("Error fetching data:", err)
		app.serverErrorResponse(w, r, err)
		return
	}
	defer res.Body.Close() // Ensure the body is closed

	// Read response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		app.serverErrorResponse(w, r, err)
		return
	}

	// Encapsulate response in envelope
	env := envelope{
		"encapsulatedData": string(data), // Convert []byte to string
	}

	// Send JSON response
	err = app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
