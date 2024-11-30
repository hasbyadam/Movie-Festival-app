package request

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type UpsertMovies struct {
	Id          uuid.NullUUID `json:"id"`
	Title       string        `json:"title"`
	Duration    null.String   `json:"duration"`
	Description null.String   `json:"description"`
	Actors      []null.String `json:"actors"`
	Genres      []null.Int    `json:"genres"`
	WatchUrl    null.String   `json:"watchUrl"`
	PosterUrl   null.String   `json:"posterUrl"`
}

type UpsertMovieViewerships struct {
	Id            uuid.UUID `json:"id"`
	MovieId       uuid.UUID `json:"movieId"`
	WatchDuration int64     `json:"watchDuration"`
}
