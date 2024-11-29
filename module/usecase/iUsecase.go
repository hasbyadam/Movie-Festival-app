package usecase

import (
	"context"
	"movie-festival-app/module/store"

	"github.com/labstack/echo/v4"
)

type Methods struct {
	Stores store.StoreInterface
}

func New(stores store.StoreInterface) UsecaseInterface {
	return &Methods{
		Stores: stores,
	}
}

type UsecaseInterface interface {
	Test(ctx context.Context)
	
	Middleware(next echo.HandlerFunc) echo.HandlerFunc
}
