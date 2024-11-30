package usecase

import (
	"context"
	"mime/multipart"
	"movie-festival-app/entity"
	"movie-festival-app/module/store"
	"movie-festival-app/schema/request"
	"movie-festival-app/schema/response"
)

type Methods struct {
	Stores store.StoreInterface
	Config *entity.Config
}

func New(stores store.StoreInterface, config *entity.Config) UsecaseInterface {
	return &Methods{
		Stores: stores,
		Config: config,
	}
}

type UsecaseInterface interface {
	UploadFile(ctx context.Context, file multipart.File, header *multipart.FileHeader, contentType string) (res response.UploadFileResponse, err error)
	UpsertMovies(ctx context.Context, req request.UpsertMovies) (err error)
	UpsertMovieViewerships(ctx context.Context, req request.UpsertMovieViewerships) (err error)
	GetMostViewedMovieAndGenre(ctx context.Context) (res response.MostViewedMovieAndGenre, err error)
}
