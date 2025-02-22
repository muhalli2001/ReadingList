// here we go!


package main
import (
	"fmt"
	//"encoding/json"
	"net/http"
	"log"
	//"time" // new import
	//"github.com/muhalli2001/ReadingList/internal/data" // New import data
	//"github.com/muhalli2001/ReadingList/internal/validator" // validator
	"io/ioutil"
)

func (app *application) searchBookHandler(w http.ResponseWriter, r *http.Request) {


	// the json should likely just include the user's search query. This should be limited to like 300 characters?
	// going to try a regular get request, see if it goes through and prints when this api gets pinged.

	res, err := http.Get("https://openlibrary.org/search.json?q=the+lord+of+the+rings")

	//check response errors:
	if err != nil {
		log.Fatal(err)
	}

	//reading the body into data a string
	data, _ := ioutil.ReadAll( res.Body )

	//into our json use:
	// actualenv := envelope{
	// 	"encapsulatedData":data,
	// }

	//close response body
	res.Body.Close()

	//print data into a string
	fmt.Printf("%s\n", data)

	env := envelope{
		"status": "available",
		"system_info": map[string]string{
		"environment": app.config.env,
		"version": version,
		},
		}

	err = app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
	// Use the new serverErrorResponse() helper.
	app.serverErrorResponse(w, r, err)
	}


}