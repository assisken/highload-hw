package app

import (
	"encoding/json"
	owm "github.com/briandowns/openweathermap"
	"net/http"
	"os"
)

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
	if err != nil {
		handleError(w, err)
		return
	}

	err = weather.CurrentByName(city)
	if err != nil {
		handleError(w, err)
		return
	}

	forecast := Forecast{city, "celsius", weather.Main.Temp}
	out, err := json.Marshal(forecast)
	if err != nil {
		handleError(w, err)
		return
	}

	_, err = w.Write(out)
	if err != nil {
		handleError(w, err)
		return
	}
}
