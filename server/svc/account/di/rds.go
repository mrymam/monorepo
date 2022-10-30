package di

import (
	"fmt"

	"github.com/onyanko-pon/monorepo/server/pkg/env"
	"github.com/onyanko-pon/monorepo/server/pkg/rds"
	"github.com/onyanko-pon/monorepo/server/svc/account/config"
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
