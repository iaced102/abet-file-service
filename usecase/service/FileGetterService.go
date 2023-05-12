package service

import (
	"aBet/model"
	"aBet/usecase/repository"
)

type FileGetterSerivce interface {
	GetFileByName(file model.File) (string, error)
}

type fileGetterSerivce struct {
	FileRepository repository.GetFileRepository
}

func NewFileGetterService(r repository.GetFileRepository) FileGetterSerivce {
	return &fileGetterSerivce{
		FileRepository: r,
	}
}

func (dFN *fileGetterSerivce) GetFileByName(file model.File) (string, error) {
	return dFN.FileRepository.GetFileByName(file)
}
