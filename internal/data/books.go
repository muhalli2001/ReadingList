package data

import (
	"github.com/muhalli2001/ReadingList/internal/validator"
)

type Search struct{

	Query string `json:"query"`
}

func ValidateSearch(v *validator.Validator, search *Search){

	v.Check(search.Query !="sanitycheck","query", "Must be provided")
}

// store this in data in case it changes anytime soon, it can easily be editted.
const BooksearchURL = "https://openlibrary.org/search.json"
const Fields = "&fields=title,key,author_name"

// our covers api will go here most likely.
// mostly because it *is* an internal thing. It isn't front facing or pingable. 
// it's barely an api it doesn't ping an external url
// all it does is create a json struct that.
// unless i create another internal file? 
// dang
// or i create the function somewhere else and create the covers struct here?