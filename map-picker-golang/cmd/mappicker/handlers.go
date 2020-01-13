package main

import (
	"fmt"
	"net/http"
)

func teams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Something happens here")
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Something happens here")

}

//func notFound(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "text/html")
//	fmt.Fprint(w, "<h1>404 - Content Not Found</h1><p>Something went wrong, the page you are looking for does not exist. Please try again.</p>")
//}
//
//var _404 http.Handler = http.HandlerFunc(notFound)
