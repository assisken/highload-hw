package app

import (
	"fmt"
	"log"
	"net/http"
)

func handleError(w http.ResponseWriter, err error) {
	message := fmt.Sprintf(`{"error": "%s"}`, err.Error())
	_, _ = w.Write([]byte(message))
	log.Println(err)
}
