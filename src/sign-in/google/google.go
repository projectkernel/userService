package google

import (
	"google.golang.org/api/plus/v1"
	"golang.org/x/oauth2"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"auth/src/pojo"
)

var config = &oauth2.Config{
	// TODO Env Vars
	ClientID:     "500595983106-7hrutl5hjnteimn0o8mmn05rskul0fhr.apps.googleusercontent.com",
	ClientSecret: "3CyXM-Y7GUpImQtMWg1klOzM",
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

func auth(authCode string) (accessToken string, refreshToken string, err error){
	token, err := config.Exchange(context.Background(), authCode)
	if err != nil {
		return "", "", err
	}
	return token.AccessToken, token.RefreshToken, nil
}

func info(accessToken string) (user pojo.UserData, err error){
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

	user = pojo.UserData{
		Name: data.DisplayName,
		Email: data.Emails[0].Value,
		ImageURL: data.Image.Url,
		Language: data.Language,
		Location: data.CurrentLocation,
		Provider: "google",
	}
	return user, nil
}