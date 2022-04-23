package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello World!")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
