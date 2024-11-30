package usecase

import (
	"context"
	"mime/multipart"
	"movie-festival-app/entity"
	"movie-festival-app/module/store"

	"github.com/labstack/echo/v4"
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
	UploadFile(ctx context.Context, file multipart.File, header *multipart.FileHeader, contentType string) (path string, err error)
	Middleware(next echo.HandlerFunc) echo.HandlerFunc
}
