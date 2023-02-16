package di

import (
	"fmt"

	"github.com/mrymam/monorepo/server/pkg/env"
	"github.com/mrymam/monorepo/server/pkg/rds"
	"github.com/mrymam/monorepo/server/svc/post/config"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {

	if env.IsProd() || env.IsDev() {
		c, err := config.GetRdsConfig()
		if err != nil {
			return nil, err
		}
		return rds.New(c)
	}
	return nil, fmt.Errorf("unknown environment.")
}
