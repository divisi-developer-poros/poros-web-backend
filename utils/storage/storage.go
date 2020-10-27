package storage

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/divisi-developer-poros/poros-web-backend/config"
	"github.com/divisi-developer-poros/poros-web-backend/utils/random"
	"github.com/gin-gonic/gin"
)

// SingleHandler mengatur upload single file
func SingleHandler(c *gin.Context) {
	filename, err := StoreFile(c, "file", config.AssetSample)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.String(http.StatusOK, `Berhasil mengupload file dengan nama:%s!\n
		Akses file anda disini: %s/%s%s`, filename, c.Request.Host, config.AssetSample, filename)
	}
}

// MultipleHandler mengatur upload multiple file
func MultipleHandler(c *gin.Context) {
	filenames, err := StoreFiles(c, "files", config.AssetSample)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.String(http.StatusOK, "Berhasil mengupload file! list nama file: %v", filenames)
	}
}

// StoreFile menyimpan satu file yang terupload ke lokasi yang dituju
func StoreFile(c *gin.Context, field string, location string) (string, error) {
	file, err := c.FormFile(field)
	if err != nil {
		return "", err
	}

	filename, err := store(c, file, location)
	if err != nil {
		return "", err
	}

	return filename, nil
}

// StoreFiles menyimpan beberapa file yang terupload ke lokasi yang dituju
func StoreFiles(c *gin.Context, field string, location string) (*[]string, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, err
	}

	files := form.File[field]

	filenames := make([]string, 0, 4)

	if len(files) == 0 {
		return nil, fmt.Errorf("file tidak ada")
	}

	for _, file := range files {
		filename, err := store(c, file, location)
		if err != nil {
			return nil, err
		}
		filenames = append(filenames, filename)
	}
	return &filenames, nil
}

// RemoveFile menghapus file yang dituju
func RemoveFile(path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}

func store(c *gin.Context, file *multipart.FileHeader, location string) (string, error) {
	filename := random.RandomString(40) + filepath.Ext(file.Filename)
	dest := location + filename

	if err := c.SaveUploadedFile(file, dest); err != nil {
		return "", err
	}
	return filename, nil
}
