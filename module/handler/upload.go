package handler

import (
	"bytes"
	"context"
	"errors"
	"io"
	"movie-festival-app/constant"
	"movie-festival-app/pkg/util"
	"movie-festival-app/schema/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) UploadMovieFile(c echo.Context) error {
	var res response.UploadFileResponse

	file, header, err := c.Request().FormFile("file")
	if err != nil {
		return err
	}
	defer file.Close()

	// validate file size
	if header.Size > constant.MaxMovieUploadSize {
		return err
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return err
	}
	defer buf.Reset()

	// validate file type
	filetype := http.DetectContentType(buf.Bytes())
	var ok bool
	for _, v := range constant.AudioVideoTypes {
		if v == filetype {
			ok = true
			break
		}
	}
	if !ok {
		return util.ErrorBadRequest(c, errors.New("invalid file type"), res)
	}

	res, err = h.Usecase.UploadFile(context.Background(), file, header, constant.Video)
	if err != nil {
		return util.ErrorInternalServerResponse(c, err, res)
	}

	return util.SuccessResponse(c, "success upload movie", res)
}

func (h *Handler) UploadImageFile(c echo.Context) error {
	var res response.UploadFileResponse

	file, header, err := c.Request().FormFile("file")
	if err != nil {
		return err
	}
	defer file.Close()

	// validate file size
	if header.Size > constant.MaxImageUploadSize {
		return err
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return err
	}
	defer buf.Reset()

	// validate file type
	filetype := http.DetectContentType(buf.Bytes())
	var ok bool
	for _, v := range constant.ImageTypes {
		if v == filetype {
			ok = true
			break
		}
	}
	if !ok {
		return util.ErrorBadRequest(c, errors.New("invalid file type"), res)
	}

	res, err = h.Usecase.UploadFile(context.Background(), file, header, constant.Image)
	if err != nil {
		return util.ErrorInternalServerResponse(c, err, res)
	}

	return util.SuccessResponse(c, "success upload image", res)
}
