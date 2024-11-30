package handler

import (
	"context"
	"movie-festival-app/pkg/util"
	"movie-festival-app/schema/request"

	"github.com/labstack/echo/v4"
)

func (h *Handler) UpsertMovies(c echo.Context) error {
	var req request.UpsertMovies

	err := c.Bind(&req)
	if err != nil {
		return util.ErrorInternalServerResponse(c, err, nil)
	}

	if err = h.Usecase.UpsertMovies(context.Background(), req); err != nil {
		return util.ErrorInternalServerResponse(c, err, nil)
	}

	return util.SuccessResponse(c, "success upsert movie", nil)
}

func (h *Handler) UpsertMovieViewerships(c echo.Context) error {
	var req request.UpsertMovieViewerships

	err := c.Bind(&req)
	if err != nil {
		return util.ErrorInternalServerResponse(c, err, nil)
	}

	if err = h.Usecase.UpsertMovieViewerships(context.Background(), req); err != nil {
		return util.ErrorInternalServerResponse(c, err, nil)
	}

	return util.SuccessResponse(c, "success upsert movie viewership", nil)
}

func (h *Handler) GetMostViewedMovieAndGenre(c echo.Context) error {
	res, err := h.Usecase.GetMostViewedMovieAndGenre(context.Background())
	if err != nil {
		return util.ErrorInternalServerResponse(c, err, res)
	}

	return util.SuccessResponse(c, "success get", res)
}
