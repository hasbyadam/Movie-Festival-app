package usecase

import (
	"context"
	"io"
	"mime/multipart"
	"movie-festival-app/constant"
	"movie-festival-app/schema/response"
	"os"
	"path/filepath"
)

func (u *Methods) UploadFile(ctx context.Context, file multipart.File, header *multipart.FileHeader, contentType string) (res response.UploadFileResponse, err error) {
	var storagePath string
	switch contentType {
	case constant.Image:
		storagePath = u.Config.Storage.BasePath + u.Config.Storage.ImagePath
	case constant.Video:
		storagePath = u.Config.Storage.BasePath + u.Config.Storage.VideoPath
	}

	path := filepath.Join(storagePath, header.Filename)

	if err = os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return
	}

	f, err := os.Create(path)
	if err != nil {
		return 
	}
	defer f.Close()

	if _, err = io.Copy(f, file); err != nil {
		return 
	}

	res.Path = path
	return 
}
