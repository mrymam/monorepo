package rdb

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteConn struct {
	filepath string
}

func (c SqliteConn) Dialector() (gorm.Dialector, error) {
	return sqlite.Open(c.filepath), nil
}
