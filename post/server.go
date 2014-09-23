package main

import (
        "log"
	"net/http"
        "encoding/json"
        "fmt"
)

func main() {
        http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/projects", newProject)
	http.HandleFunc("/projects2", newProject2)
	http.ListenAndServe(":8080", nil)
}

// post 2 integers
func newProject(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, r.Method+" not allowed", http.StatusMethodNotAllowed)
		return
	}
	req := struct{ A, B int }{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	res := struct{ Result int }{req.A}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "oops", http.StatusInternalServerError)
	}
}

// post 2 strings
func newProject2(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, r.Method+" not allowed", http.StatusMethodNotAllowed)
		return
	}

        fmt.Printf("%+v\n", r.Body)

	req := struct{ A, B string }{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
                fmt.Println("wow. error")
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
        fmt.Println("life is good")
	res := struct{ Result string }{req.A}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "oops", http.StatusInternalServerError)
	}
}
