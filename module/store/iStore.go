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
	Test(ctx context.Context)
}
