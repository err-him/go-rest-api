package config

import (
	"go-rest-api/src/controller"

	"github.com/gorilla/mux"
)

func handleAppRoutes(r *mux.Router) {

	//handling API versioning
	v1 := r.PathPrefix("/api/v1").Subrouter()

	v1.HandleFunc("/", controller.GetNumber).Methods("GET")
	v1.HandleFunc("/post", controller.PostNumber).Methods("POST")

}
