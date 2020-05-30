package main

import (
	"github.com/lacledeslan/map_picker2/pkg/routes"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
	r := mux.NewRouter()

	http.Handle("/", r)
	routes.RegisterGameRoutes(r)
	routes.RegisterMiscRoutes(r)
	//r.HandleFunc("/teams", teams)
	//r.HandleFunc("/users", users)
	http.ListenAndServe(":8080", r)
}
