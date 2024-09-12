package main

import (
	"net/http"
)

func main() {
	// hello()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Golang Server!"))
	})
	http.ListenAndServe(":8080", nil)
}
