package di

import (
	"time"

	"github.com/mrymam/monorepo/server/pkg/setting"
	"github.com/mrymam/monorepo/server/svc/authn/domain/svc"
	svcImple "github.com/mrymam/monorepo/server/svc/authn/svc"
)

func GetTokenSvc() (svc.TokenSvc, error) {
	return svcImple.InitJwtSvc(setting.Get().JWT.SecretKey, time.Hour*24*30)
}
