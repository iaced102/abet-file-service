package initapp

import (
	"aBet/adapters/repository"
	"aBet/infrastructure/database/connection"
	"aBet/infrastructure/router"
	"aBet/registry"
	"fmt"
	"log"
	"os"

	sLog "aBet/log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func InitApp(envPath string) (*echo.Echo, *repository.Orm) {
	godotenv.Load(envPath)
	db, err := connection.NewPostgresCon().Conn()
	fmt.Println(db, err)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.CloseDB()
	sLog.NewLogger()
	sLog.Info("Server listen at http://localhost"+":"+os.Getenv("SERVER_PORT"), map[string]interface{}{"line": sLog.Trace()})

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())
	if err := e.Start("0.0.0.0:" + os.Getenv("SERVER_PORT")); err != nil {
		log.Fatalln(err)
	}

	return e, db
}
