package google

import (
	"google.golang.org/api/plus/v1"
	"golang.org/x/oauth2"
	"golang.org/x/net/context"
	"auth/src/pojo"
)

func (Provider) Info(accessToken string) (user *pojo.User, err error) {
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