package middleware

import (
	"movie-festival-app/entity"
	"movie-festival-app/module/store"
)

type Middleware struct {
	config *entity.Config
	store  store.StoreInterface
}

func NewMiddleware(config *entity.Config, store store.StoreInterface) *Middleware {
	return &Middleware{
		store:  store,
		config: config,
	}
}
