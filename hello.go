package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if message := os.Getenv("MESSAGE"); message != "" {
			w.Write([]byte(message))
		} else {
			w.Write([]byte("Hello, world!\n"))
		}
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
