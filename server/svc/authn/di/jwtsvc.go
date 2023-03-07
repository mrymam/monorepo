package di

import (
	"time"

	"github.com/mrymam/radio_rec/server/pkg/setting"
	"github.com/mrymam/radio_rec/server/svc/authn/domain/svc"
	svcImple "github.com/mrymam/radio_rec/server/svc/authn/svc"
)

func GetTokenSvc() (svc.TokenSvc, error) {
	return svcImple.InitJwtSvc(setting.Get().JWT.SecretKey, time.Hour*24*30)
}
