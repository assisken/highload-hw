package main

import (
	"log"
	"net/http"
)
import "highload-hw/app"

func main() {
	port := ":8000"
	log.Printf("Starting listening on port %s", port)
	err := http.ListenAndServe(port, app.Handlers())
	if err != nil {
		log.Fatal(err)
	}
}