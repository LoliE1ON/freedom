package utils

import (
	"io"
	"mime/multipart"
	"os"
)

func UploadFile(file *multipart.FileHeader, savePath string) (err error) {

	src, err := file.Open()
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(savePath)
	if err != nil {
		return
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return
	}

	return err
}
