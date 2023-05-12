package registry

import (
	"aBet/adapters/controller"
	"aBet/adapters/repository"
)

type registry struct {
	db *repository.Orm
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *repository.Orm) Registry {
	return &registry{
		db: db,
	}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		FileController: r.NewFileController(),
	}
}
