package connection

import (
	"aBet/adapters/repository"
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
	_ "github.com/lib/pq"
)

// A PostgresCon represents a mysql connection
type PostgresCon struct{}

// NewPostgresCon creates a PostgresCon struct.
func NewPostgresCon() PostgresConInterface {
	return &PostgresCon{}
}

// Conn get a mysql connection.
func (mc *PostgresCon) Conn() (*repository.Orm, error) {
	// dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", , , , , )
	// conn, err := sql.Open(os.Getenv("DB_DRIVER"), dataSourceName)
	conn := pg.Connect(&pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
		Addr:     os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
	})
	fmt.Println(conn)
	return repository.NewOrm(conn), nil
}
