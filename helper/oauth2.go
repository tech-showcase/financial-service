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
	oauth2Helper struct {
		config oauth2.Config
	}
	OAuth2Helper interface {
		HandleAuthorizeRequest() (result string)
		HandleTokenRequest(authorizationCode string) (token Token, err error)
	}
)

var (
	OAuth2HelperInstance OAuth2Helper
)

const (
	serviceAddress = "http://localhost:8082"
)

func NewOAuth2Helper(authServiceAddress string) OAuth2Helper {
	var instance oauth2Helper
	instance.config = newOAuth2Config(authServiceAddress)

	return &instance
}

func newOAuth2Config(authServiceAddress string) oauth2.Config {
	authorizeEndpoint, _ := JoinURL(authServiceAddress, "/oauth2/authorize")
	tokenEndpoint, _ := JoinURL(authServiceAddress, "/oauth2/token")
	redirectEndpoint, _ := JoinURL(serviceAddress, "/oauth2/token")

	return oauth2.Config{
		ClientID:     "082222333444",
		ClientSecret: "082222333444",
		Scopes:       []string{"all"},
		RedirectURL:  redirectEndpoint,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizeEndpoint,
			TokenURL: tokenEndpoint,
		},
	}
}

func (instance *oauth2Helper) HandleAuthorizeRequest() (result string) {
	result = instance.config.AuthCodeURL("xyz")
	return
}

func (instance *oauth2Helper) HandleTokenRequest(authorizationCode string) (token Token, err error) {
	oAuth2Token, err := instance.config.Exchange(context.Background(), authorizationCode)
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
