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
	http.HandleFunc("/projects3", newProject3)
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
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	res := struct{ Result int }{req.A}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "oops", http.StatusInternalServerError)
	}
}

// endpoint for form POST
func newProject2(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, r.Method+" not allowed", http.StatusMethodNotAllowed)
		return
	}

        str := fmt.Sprintf("%s %s", r.FormValue("A"),  r.FormValue("B"))
        log.Println(str)

        w.WriteHeader(http.StatusCreated)
}

// endpoint for JSON POST
func newProject3(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, r.Method+" not allowed", http.StatusMethodNotAllowed)
		return
	}

	req := struct{ A, B string }{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	res := struct{ Result string }{req.A}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "oops", http.StatusInternalServerError)
	}
        w.WriteHeader(http.StatusCreated)
}
