package data
import (
 "time"
 "github.com/muhalli2001/ReadingList/internal/validator"
)
type Movie struct {
ID int64 `json:"id"`
CreatedAt time.Time `json:"-"`
Title string `json:"title"`
Year int32 `json:"year,omitempty"`
// Use the Runtime type instead of int32. Note that the omitempty directive will
// still work on this: if the Runtime field has the underlying value 0, then it will
// be considered empty and omitted -- and the MarshalJSON() method we just made
// won't be called at all.
Runtime Runtime `json:"runtime,omitempty"`
Genres []string `json:"genres,omitempty"`
Version int32 `json:"version"`
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
    v.Check(movie.Title != "", "title", "must be provided")
    v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")
    v.Check(movie.Year != 0, "year", "must be provided")
    v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
    v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")
    v.Check(movie.Runtime != 0, "runtime", "must be provided")
    v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")
    v.Check(movie.Genres != nil, "genres", "must be provided")
    v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
    v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
    v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
    }

// adding some stuff as a test