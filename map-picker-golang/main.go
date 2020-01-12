package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func games(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Something happens here")
}

func teams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Something happens here")
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Something happens here")

}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>404 - Content Not Found</h1><p>Something went wrong, the page you are looking for does not exist. Please try again.</p>")
}

var _404 http.Handler = http.HandlerFunc(notFound)

func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = _404
	r.HandleFunc("/games", games)
	r.HandleFunc("/teams", teams)
	r.HandleFunc("/users", users)
	http.ListenAndServe(":8080", r)
}
