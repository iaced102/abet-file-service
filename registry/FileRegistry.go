package registry

import (
	"aBet/adapters/controller"
	uR "aBet/adapters/repository"

	// "aBet/usecase/repository"
	uSv "aBet/usecase/service"
)

func (r *registry) NewFileController() controller.FileController {
	return controller.NewFileController(
		r.NewAddNewFileService(),
		r.FileGetterService(),
	)
}

// fileUPloaderService: upload file
func (r *registry) NewAddNewFileService() uSv.FileUploaderService {
	return uSv.NewFileUploaderService(uR.NewFileRepository(r.db))
}

func (r *registry) FileGetterService() uSv.FileGetterSerivce {
	return uSv.NewFileGetterService(uR.NewFileRepository(r.db))
}
