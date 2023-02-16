package repository

import (
	"encoding/json"
	"io/ioutil"

	"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter"
	"github.com/mrymam/monorepo/server/pkg/setting"
	"github.com/mrymam/monorepo/server/svc/authn/domain/model"
	"github.com/mrymam/monorepo/server/svc/authn/infra/entity"
)

func InitTwitterUserRepo() (TwitterUserRepo, error) {
	s := setting.Get().Twitter.OAuth1
	return TwitterUserRepo{
		cfg: oauth1.Config{
			ConsumerKey:    s.ConsumerKey,
			ConsumerSecret: s.ConsumerSecret,
			CallbackURL:    s.CallbackURL,
			Endpoint:       twitter.AuthorizeEndpoint,
		},
	}, nil
}

type TwitterUserRepo struct {
	cfg oauth1.Config
}

func (r TwitterUserRepo) Get(token *oauth1.Token) (model.TwitterUser, error) {
	clt := r.cfg.Client(oauth1.NoContext, token)

	path := "https://api.twitter.com/1.1/account/verify_credentials.json"
	res, err := clt.Get(path)
	if err != nil {
		return model.TwitterUser{}, err
	}

	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return model.TwitterUser{}, err
	}
	var e entity.TwitterUser
	err = json.Unmarshal([]byte(b), &e)
	if err != nil {
		return model.TwitterUser{}, err
	}

	return e.ToModel(), nil
}
