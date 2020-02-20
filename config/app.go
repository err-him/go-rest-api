package config

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//App has Router instances

type App struct {
	Router *mux.Router
}

//Initialize app with the Router
func (a *App) Intialize() {
	a.Router = mux.NewRouter()
	handleAppRoutes(a.Router)

}

//Run App Here

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(":"+host, a.Router))
}
