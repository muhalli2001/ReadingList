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
