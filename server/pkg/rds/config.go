package rds

import (
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
