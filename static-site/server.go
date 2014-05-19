package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
        http.HandleFunc("/static/", static)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func static(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, r.URL.Path[1:])
}
