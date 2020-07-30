package presenter

import (
	"context"
	"encoding/json"
	"golang.org/x/oauth2"
	"net/http"
)

const (
	authServerURL = "http://localhost:8080/oauth2"
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
	oauth2Token *oauth2.Token
)

func Authorize(w http.ResponseWriter, r *http.Request) {
	u := oauth2Config.AuthCodeURL("xyz")
	http.Redirect(w, r, u, http.StatusFound)
}

func Token(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	state := r.Form.Get("state")
	if state != "xyz" {
		http.Error(w, "State invalid", http.StatusBadRequest)
		return
	}

	code := r.Form.Get("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}

	token, err := oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	oauth2Token = token

	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(token)
}
