// hello world using anonymous function

package main

import (
        "io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                io.WriteString(w, "hello, world!\n")
        })

	http.ListenAndServe(":3000", nil)
}
