package repository

import (
	"github.com/go-pg/pg/v10"
)

type Orm struct {
	pgdb *pg.DB
}

func (sorm *Orm) CloseDB() error {
	return sorm.pgdb.Close()
}

func NewOrm(pgdb *pg.DB) *Orm {
	return &Orm{
		pgdb: pgdb,
	}
}
