package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
}

func NewApiServer() *ApiServer {
	return &ApiServer{}
}

func (as *ApiServer) Run() {
	r := mux.NewRouter()
	r.HandleFunc("", ServHome).Methods("GET")
}

func ServHome(w http.ResponseWriter, r *http.Request) {

}
