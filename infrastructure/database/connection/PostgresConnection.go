package connection

import "aBet/adapters/repository"

// A MySQLConn represents a database connection.
type PostgresConInterface interface {
	Conn() (*repository.Orm, error)
}
