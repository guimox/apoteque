package api

import (
	"net/http"
)

func (app *application) createLinks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Links"))
}
