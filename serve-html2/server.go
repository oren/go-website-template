package main

import (
        "fmt"
	"net/http"
        "io/ioutil"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	dat, err := ioutil.ReadFile("index.html")
	check(err)
	fmt.Fprintf(w, string(dat))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
