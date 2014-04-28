// hello world using a named function - handler

package main

import (
        "io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
