package repository

import (
	"aBet/model"
	"mime/multipart"
)

type AddNewFileRepository interface {
	AddNewFile(file *model.File, uploadFile *multipart.FileHeader) (*model.File, error)
}

type GetFileRepository interface {
	GetFileByName(file model.File) (string, error)
}
