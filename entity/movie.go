package entity

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Movie struct {
	Id          uuid.UUID
	Title       string
	Duration    null.String
	Description null.String
	Actors      []null.String
	WatchUrl    null.String
	CreatedAt   int64
	UpdatedAt   null.Int
	PosterUrl   null.String
}

type MovieViewerships struct {
	Id            uuid.UUID
	MovieId       uuid.UUID
	WatchDuration int64
	CreatedAt     int64
	UpdatedAt     null.Int
}

type MovieVotes struct {
	Id        uuid.UUID
	MovieId   uuid.UUID
	UserId    uuid.UUID
	CreatedAt int64
}

type MovieGenres struct {
	Id        uuid.UUID
	MovieId   uuid.UUID
	GenreId   int64
	CreatedAt int64
}
