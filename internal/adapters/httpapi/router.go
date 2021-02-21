package httpapi

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *St) router() http.Handler {
	r := mux.NewRouter()

	r.PathPrefix("/send").HandlerFunc(a.hSend).Methods("POST")

	return r
}
