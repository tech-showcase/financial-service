package presenter

import (
	"encoding/json"
	"github.com/tech-showcase/financial-service/controller"
	"net/http"
)

func Authorize(w http.ResponseWriter, r *http.Request) {
	u := controller.Authorize()
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

	token, err := controller.Token(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(token)
}
