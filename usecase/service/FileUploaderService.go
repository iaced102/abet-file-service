package service

import (
	"aBet/model"
	"aBet/usecase/repository"
	"mime/multipart"
)

type FileUploaderService interface {
	AddNewFile(file *model.File, uploadFile *multipart.FileHeader) (*model.File, error)
}

type fileUploaderService struct {
	FileRepository repository.AddNewFileRepository
}

func NewFileUploaderService(r repository.AddNewFileRepository) FileUploaderService {
	return &fileUploaderService{
		FileRepository: r,
	}
}

func (fS fileUploaderService) AddNewFile(file *model.File, uploadFile *multipart.FileHeader) (*model.File, error) {
	return fS.FileRepository.AddNewFile(file, uploadFile)
}
