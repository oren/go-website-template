package main

import (
        "log"
	"net/http"
        "encoding/json"
)

func main() {
        http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/projects", newProject)
	http.ListenAndServe(":8080", nil)
}

func newProject(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, r.Method+" not allowed", http.StatusMethodNotAllowed)
		return
	}
	req := struct{ A, B string }{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	res := struct{ Result string }{req.A}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "oops", http.StatusInternalServerError)
	}
}
