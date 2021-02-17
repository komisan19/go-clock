package main

import (
	"log"
	"net/http"
)

func main() {
	port := "8080"
	http.Handle("/", http.FileServer(http.Dir("./docs/")))
	log.Printf("Listen on port: %s", port)
	http.ListenAndServe(":"+port, nil)
}
