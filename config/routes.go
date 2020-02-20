package config

import (
	"go-rest-api/src/controller"

	"github.com/gorilla/mux"
)

func handleAppRoutes(r *mux.Router) {
	r.HandleFunc("/", controller.GetNumber).Methods("GET")
	r.HandleFunc("/post", controller.PostNumber).Methods("POST")

}
