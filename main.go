package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	os.Mkdir("tmp", os.ModePerm) // Test create a folder
	os.Create("tmp/tmp")         // Test create a file

	log.Println("The app starts at 8080.")
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello World!")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
