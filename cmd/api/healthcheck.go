// JSON with and without marshalling

// package main
// import (
// "fmt"
// "net/http"
// )
// func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
// // Create a fixed-format JSON response from a string. Notice how we're using a raw
// // string literal (enclosed with backticks) so that we can include double-quote
// // characters in the JSON without needing to escape them? We also use the %q verb to
// // wrap the interpolated values in double-quotes.
// js := `{"status": "available", "environment": %q, "version": %q}`
// js = fmt.Sprintf(js, app.config.env, version)
// // Set the "Content-Type: application/json" header on the response. If you forget to
// // this, Go will default to sending a "Content-Type: text/plain; charset=utf-8"
// // header instead.
// w.Header().Set("Content-Type", "application/json")
// // Write the JSON as the HTTP response body.
// w.Write([]byte(js))
// }

// writing json using marshal code

// package main
// import (
// "encoding/json" // New import
// "net/http"
// )
// func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
// // Create a map which holds the information that we want to send in the response.
// data := map[string]string{
// "status": "available",
// "environment": app.config.env,
// "version": version,
// }
// // Pass the map to the json.Marshal() function. This returns a []byte slice
// // containing the encoded JSON. If there was an error, we log it and send the client
// // a generic error message.
// js, err := json.Marshal(data) // marshall a map

// if err != nil {
// app.logger.Println(err)
// http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
// return
// }

// // Append a newline to the JSON. This is just a small nicety to make it easier to
// // view in terminal applications.
// js = append(js, '\n')
// // At this point we know that encoding the data worked without any problems, so we
// // can safely set any necessary HTTP headers for a successful response.
// w.Header().Set("Content-Type", "application/json")
// // Use w.Write() to send the []byte slice containing the JSON as the response body.
// w.Write(js)
// }

// moving the json writing functionality to a function in helpers.go cus its a helper fr

package main
import (
"net/http"
)
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
data := map[string]string{
"status": "available",
"environment": app.config.env,
"version": version,
}
err := app.writeJSON(w, http.StatusOK, data, nil)
if err != nil {
app.logger.Println(err)
http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
}
}