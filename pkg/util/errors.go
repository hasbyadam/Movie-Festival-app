package util

import (
	"movie-festival-app/schema/response"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func SuccessResponse(c echo.Context, message string, data interface{}) error {
	responseData := response.Base{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    message,
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}
	return c.JSON(http.StatusOK, responseData)
}

// ErrorInternalServerResponse returns
func ErrorInternalServerResponse(ctx echo.Context, err error, data interface{}) error {
	responseData := response.Base{
		Status:     "internal server error",
		StatusCode: http.StatusInternalServerError,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	return ctx.JSON(http.StatusInternalServerError, responseData)
}

// ErrorBadRequest returns
func ErrorBadRequest(ctx echo.Context, err error, data interface{}) error {
	responseData := response.Base{
		Status:     "bad request",
		StatusCode: http.StatusBadRequest,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	return ctx.JSON(http.StatusBadRequest, responseData)
}
