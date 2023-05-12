package router

/*
create by: Hoangnd
create at: 2023-01-01
des: Xử lý router & authen
*/

import (
	aAuth "aBet/adapters/auth"
	"aBet/adapters/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	var authObject aAuth.AuthObject
	// crypt.CreateAndSaveKeyPair()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ct := &controller.Context{
				Context: c,
			}
			return next(ct.Context)
		}
	})

	e.Static("/", "static/index.html")
	e.POST("/uploadS", func(context echo.Context) error { return forward(context, authObject, c.FileController.AddNewFile) })
	e.GET("/readFile/:fileName", func(context echo.Context) error { return forward(context, authObject, c.FileController.ReadFile) })

	return e
}
func forward(context echo.Context, authObject aAuth.AuthObject, f func(*controller.Context) error) error {
	ct := &controller.Context{}
	ct.Context = context
	ct.AuthObject = authObject
	return f(ct)
}
