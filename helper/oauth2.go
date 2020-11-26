package helper

import (
	"context"
	"golang.org/x/oauth2"
	"time"
)

type (
	Token struct {
		AccessToken  string    `json:"access_token"`
		TokenType    string    `json:"token_type,omitempty"`
		RefreshToken string    `json:"refresh_token,omitempty"`
		Expiry       time.Time `json:"expiry,omitempty"`
	}
	oauth2Helper struct{}
	OAuth2Helper interface {
		HandleAuthorizeRequest() (result string)
		HandleTokenRequest(authorizationCode string) (token Token, err error)
	}
)

var (
	oauth2Config = oauth2.Config{
		ClientID:     "082222333444",
		ClientSecret: "082222333444",
		Scopes:       []string{"all"},
		RedirectURL:  "http://localhost:8082/oauth2/token",
		Endpoint: oauth2.Endpoint{
			AuthURL:  authServerURL + "/authorize",
			TokenURL: authServerURL + "/token",
		},
	}

	OAuth2HelperInstance OAuth2Helper
)

const (
	authServerURL = "http://localhost:8080/oauth2"
)

func NewOAuth2Helper() OAuth2Helper {
	var instance oauth2Helper

	return &instance
}

func (instance *oauth2Helper) HandleAuthorizeRequest() (result string) {
	result = oauth2Config.AuthCodeURL("xyz")
	return
}

func (instance *oauth2Helper) HandleTokenRequest(authorizationCode string) (token Token, err error) {
	oAuth2Token, err := oauth2Config.Exchange(context.Background(), authorizationCode)
	if err != nil {
		return
	}

	token = Token{
		AccessToken:  oAuth2Token.AccessToken,
		TokenType:    oAuth2Token.TokenType,
		RefreshToken: oAuth2Token.RefreshToken,
		Expiry:       oAuth2Token.Expiry,
	}
	return
}
