package rds

import (
	"gorm.io/gorm"
)

func New(c Config) (*gorm.DB, error) {
	dlc, err := c.Conn.Dialector()
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(dlc)
	if err != nil {
		return nil, err
	}
	return db, nil
}
