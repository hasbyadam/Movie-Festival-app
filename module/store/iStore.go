package store

import (
	"context"
	"database/sql"
	"movie-festival-app/connection"
	"movie-festival-app/entity"
)

type Client struct {
	postgre *sql.DB
}

func New(config *entity.Config) StoreInterface {
	return &Client{
		postgre: connection.Postgres(config.Database.Pg),
	}
}

type StoreInterface interface {
	UpsertMovies(ctx context.Context, req entity.Movie) (err error)
	InsertMovieGenres(ctx context.Context, req []entity.MovieGenres) (err error)
	GetMostViewedMovieAndGenre(ctx context.Context) (res entity.MostViewedMovieAndGenre, err error)
	UpsertMovieViewerships(ctx context.Context, req entity.MovieViewerships) (err error)
}
