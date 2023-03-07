package config

import (
	"github.com/mrymam/radio_rec/server/pkg/rds"
	"github.com/mrymam/radio_rec/server/pkg/setting"
)

func GetRdsConfig() (rds.Config, error) {
	s := setting.Get()
	c := rds.Config{
		DBMS: rds.DBMS("sqlite"),
		Conn: rds.SqliteConn{
			Filepath: s.Sqlite.Filepath,
		},
	}
	return c, nil
}
