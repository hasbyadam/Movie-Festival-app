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

	e.GET("/test", handler.Test, handler.Usecase.Middleware)

}
