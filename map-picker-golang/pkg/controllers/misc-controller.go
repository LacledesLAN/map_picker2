package controllers

import (
	"fmt"
	"net/http"
)


func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>404 - Content Not Found</h1><p>Something went wrong, the page you are looking for does not exist. Please try again.</p>")
}


var HTTP404 http.Handler = http.HandlerFunc(notFound)
