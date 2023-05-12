package controller

import (
	"net/http"
)

type AppController struct {
	FileController FileController
}

func (app *AppController) HelthCheck(e *Context) error {
	return e.Output(http.StatusOK, map[string]interface{}{"status": "Running"}, nil)
}
