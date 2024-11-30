package handler

import (
	"movie-festival-app/module/usecase"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Usecase usecase.UsecaseInterface
}

func New(e *echo.Group, usecase usecase.UsecaseInterface) {
	handler := &Handler{
		Usecase: usecase,
	}

	admin := e.Group("/admin")

	upload := admin.Group("/upload")
	upload.POST("/image", handler.UploadImageFile)
	upload.POST("/movie", handler.UploadMovieFile)

	adminMovie := admin.Group("/movie")
	adminMovie.POST("", handler.UpsertMovies)
	adminMovie.GET("/most-viewed", handler.GetMostViewedMovieAndGenre)
}
