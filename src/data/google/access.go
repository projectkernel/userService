package google

import (
	"golang.org/x/oauth2"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
)

type Provider struct {
	config *oauth2.Config
}

func NewProvider(clientId string, clientSecret string) *Provider {
	config := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:   []string{
			"https://www.googleapis.com/auth/plus.login",
			"https://www.googleapis.com/auth/plus.me",
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
		// Use "postmessage" for the code-flow for server side apps
		RedirectURL: "postmessage",
	}
	return &Provider{
		config: config,
	}
}

func (prov Provider) Exchange(authCode string) (accessToken string, refreshToken string, err error) {
	token, err := prov.config.Exchange(context.Background(), authCode)
	if err != nil {
		return "", "", err
	}
	return token.AccessToken, token.RefreshToken, nil
}