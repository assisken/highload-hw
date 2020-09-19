package app

import "net/http"

func Handlers() http.Handler {
	handler := http.NewServeMux()

	handler.HandleFunc("/v1/current/", GetCurrentForecast)

	return handler
}
