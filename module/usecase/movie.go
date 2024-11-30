package usecase

import (
	"context"
	"movie-festival-app/entity"
	"movie-festival-app/schema/request"
	"movie-festival-app/schema/response"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (u *Methods) UpsertMovies(ctx context.Context, req request.UpsertMovies) (err error) {
	var id uuid.UUID = req.Id.UUID
	if !req.Id.Valid {
		id = uuid.New()
	}

	if err = u.Stores.UpsertMovies(ctx, entity.Movie{
		Id:          id,
		Title:       req.Title,
		Duration:    req.Duration,
		Description: req.Description,
		Actors:      req.Actors,
		WatchUrl:    req.WatchUrl,
	}); err != nil {
		zap.S().Info(err)
		return
	}

	if len(req.Genres) > 0 {
		genres := []entity.MovieGenres{}
		for _, v := range req.Genres {
			genres = append(genres, entity.MovieGenres{
				MovieId: id,
				GenreId: v.Int64,
			})
		}
		if err = u.Stores.InsertMovieGenres(ctx, genres); err != nil {
			zap.S().Info(err)
		}
	}

	return
}

func (u *Methods) UpsertMovieViewerships(ctx context.Context, req request.UpsertMovieViewerships) (err error) {

	if err = u.Stores.UpsertMovieViewerships(ctx, entity.MovieViewerships{
		Id:            req.Id,
		MovieId:       req.MovieId,
		WatchDuration: req.WatchDuration,
	}); err != nil {
		zap.S().Info(err)
	}

	return
}

func (u *Methods) GetMostViewedMovieAndGenre(ctx context.Context) (res response.MostViewedMovieAndGenre, err error) {

	data, err := u.Stores.GetMostViewedMovieAndGenre(ctx)
	if err != nil {
		zap.S().Info(err)
		return
	}

	res.Movie = data.Movie
	res.Genre = data.Genre

	return
}

func (u *Methods) GetMoviesPublic(ctx context.Context, req request.GetMovies) (res response.GetMovies, err error) {

	res, err = u.Stores.GetMoviesPublic(ctx, req)
	if err != nil {
		zap.S().Info(err)
	}

	return

}
