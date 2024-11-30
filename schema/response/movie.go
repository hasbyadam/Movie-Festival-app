package response

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type MostViewedMovieAndGenre struct {
	Movie null.String `json:"movie"`
	Genre null.String `json:"genre"`
}

type GetMovies struct {
	Movies     []Movies   `json:"movies"`
	Pagination Pagination `json:"pagination"`
}

type Movies struct {
	Id          uuid.UUID     `json:"id"`
	Title       string        `json:"title"`
	Duration    null.String   `json:"duration"`
	Description null.String   `json:"description"`
	Actors      []null.String `json:"actors"`
	Genres      []null.String `json:"genres"`
	WatchUrl    null.String   `json:"watchUrl"`
	PosterUrl   null.String   `json:"posterUrl"`
}
