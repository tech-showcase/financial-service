package controller

import (
	"github.com/tech-showcase/financial-service/helper"
)

func Authorize() (result string) {
	result = helper.OAuth2HelperInstance.HandleAuthorizeRequest()
	return
}

func Token(authorizationCode string) (token helper.Token, err error) {
	token, err = helper.OAuth2HelperInstance.HandleTokenRequest(authorizationCode)
	return
}
