package oauth

import (
	"github.com/ramailh/auth-server/props"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var Conf oauth2.Config

func UrlOauth() string {
	Conf = oauth2.Config{
		ClientID:     props.ClientID,
		ClientSecret: props.ClientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		RedirectURL:  "http://localhost:9192/auth/oauth/google-callback",
	}

	endpoint := Conf.AuthCodeURL("state")

	return endpoint
}
