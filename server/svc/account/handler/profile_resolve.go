package handler

import (
	"github.com/mrymam/monorepo/server/svc/account/domain/model/profile"
)

type Profile struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func resolveProfile(p profile.Profile) (Profile, error) {
	return Profile{
		ID:   string(p.ID),
		Name: string(p.Name),
	}, nil
}

func (p Profile) ToModel() (profile.Profile, error) {
	return profile.Init(profile.Name(p.Name))
}
