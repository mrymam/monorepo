package rds

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteConn struct {
	Filepath string
}

func (c SqliteConn) Dialector() (gorm.Dialector, error) {
	return sqlite.Open(c.Filepath), nil
}
