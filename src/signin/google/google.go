package google

import (
	"google.golang.org/api/plus/v1"
	"golang.org/x/oauth2"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"auth/src/pojo"
)

type provider struct {
	config *oauth2.Config
}

func New(clientId string, clientSecret string) *provider {
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
	return &provider{
		config: config,
	}
}

func (provider provider) Auth(authCode string) (accessToken string, refreshToken string, err error){
	token, err := provider.config.Exchange(context.Background(), authCode)
	if err != nil {
		return "", "", err
	}
	return token.AccessToken, token.RefreshToken, nil
}

func (provider) Info(accessToken string) (user *pojo.User, err error){
	// Create a new authorized API client
	token := new(oauth2.Token)
	token.AccessToken = accessToken
	client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))
	service, err := plus.New(client)
	if err != nil {
		return user, err
	}

	data, err := service.People.Get("me").Do()
	if err != nil {
		return user, err
	}

	user = &pojo.User{
		Name: data.DisplayName,
		Email: data.Emails[0].Value,
		ImageURL: data.Image.Url,
		Language: data.Language,
		Location: data.CurrentLocation,
		Provider: "google",
	}
	return user, nil
}