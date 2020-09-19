package app

import (
	"encoding/json"
	owm "github.com/briandowns/openweathermap"
	"log"
	"net/http"
	"os"
)

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}
}

func GetCurrentForecast(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	city := query.Get("city")
	if city == "" {
		_, err := w.Write([]byte(`{"error": "city field is required"}`))
		handleError(w, err)
		return
	}

	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	weather, err := owm.NewCurrent("C", "EN", apiKey)
	handleError(w, err)

	err = weather.CurrentByName(city)
	handleError(w, err)

	forecast := Forecast{city, "celsius", weather.Main.Temp}
	out, err := json.Marshal(forecast)
	handleError(w, err)

	_, err = w.Write(out)
	handleError(w, err)
}
