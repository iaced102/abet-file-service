package repository

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"aBet/model"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type FileRepository interface {
	AddNewFile(file *model.File, uploadFile *multipart.FileHeader) (*model.File, error)
	GetFileByName(file model.File) (string, error)
}

type fileRepository struct {
	db *Orm
}

func NewFileRepository(db *Orm) *fileRepository {
	return &fileRepository{
		db: db,
	}
}

func (fR *fileRepository) AddNewFile(file *model.File, uploadFile *multipart.FileHeader) (*model.File, error) {
	fileName := file.Name
	er := os.Remove(fmt.Sprint("uploads", "/", fileName+".png"))
	if er == nil {
		return file, nil
	}
	er = os.Remove(fmt.Sprint("uploads", "/", fileName+".jpg"))
	if er == nil {
		return file, nil
	}
	er = os.Remove(fmt.Sprint("uploads", "/", fileName+".jepg"))
	if er == nil {
		return file, nil
	}
	localPath := "uploads/"

	uploadFileContent, e := uploadFile.Open()
	if e != nil {
		return file, e
	}
	uploadFileContent.Close()
	fullFileName := file.Name

	dst, err := os.Create(fmt.Sprint(localPath, "/", fullFileName))
	defer dst.Close()
	if err != nil {
		return file, err
	}
	if _, err = io.Copy(dst, uploadFileContent); err != nil {
		return file, err
	}
	return file, nil
}

func (fR *fileRepository) GetFileByName(file model.File) (string, error) {
	fileName := file.Name

	localFile, er := os.OpenFile(fmt.Sprint("uploads", "/", fileName+".png"), os.O_RDWR, 0644)
	localFile.Close()
	if er == nil {
		return fmt.Sprint("uploads", "/", fileName, ".png"), nil
	}
	localFile, er = os.OpenFile(fmt.Sprint("uploads", "/", fileName+".jpg"), os.O_RDWR, 0644)
	localFile.Close()
	if er == nil {
		return fmt.Sprint("uploads", "/", fileName, ".jpg"), nil
	}
	localFile, er = os.OpenFile(fmt.Sprint("uploads", "/", fileName+".jepg"), os.O_RDWR, 0644)
	localFile.Close()
	if er == nil {
		return fmt.Sprint("uploads", "/", fileName, ".jepg"), nil
	}

	return "", errors.New("notfound")
}
