package rdb

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	DBMS DBMS
	Conn Conn
}

type DBMS string

const (
	SQLite DBMS = "sqlite"
)

type Conn interface {
	Dialector() (gorm.Dialector, error)
}

type SqliteConn struct {
	filepath string
}

func (c SqliteConn) Dialector() (gorm.Dialector, error) {
	return sqlite.Open(c.filepath), nil
}
