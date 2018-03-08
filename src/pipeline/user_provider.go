package pipeline

import (
	"auth/src/data"
	"auth/src/pojo"
)

type Provider struct {
	prov      data.SocialProvider
}

func NewProvider(prov data.SocialProvider) *Provider {
	return &Provider{
		prov:      prov,
	}
}

func (prov Provider) GetUserFromProvider(token string) (user *pojo.User, err error) {
	access, refresh, err := prov.prov.Exchange(token)
	if err != nil {
		return user, err
	}
	user, err = prov.prov.Info(access)
	if err != nil {
		return user, err
	}
	user.RefreshToken = refresh
	return user, nil
}