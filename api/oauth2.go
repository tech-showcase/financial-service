package api

import (
	"github.com/tech-showcase/financial-service/presenter"
	"net/http"
)

func RegisterOAuth2HTTPAPI(mux *http.ServeMux) {
	mux.HandleFunc("/oauth2/authorize", presenter.Authorize)
	mux.HandleFunc("/oauth2/token", presenter.Token)
}
