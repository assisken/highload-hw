package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func makeResponse(city, key string, dt int64) (float64, error) {
	url := "https://api.openweathermap.org/data/2.5/forecast?q=%s&appid=%s&units=%s"
	response, err := http.Get(fmt.Sprintf(url, city, key, "celsius"))
	if response == nil {
		return .0, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return .0, err
	}

	path := fmt.Sprintf("list.#(dt==%d).main.temp", dt)
	result := gjson.Get(string(body), path)
	if !result.Exists() {
		return .0, errors.New("temperature is not found")
	}
	return result.Float(), nil
}

func GetForecast(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	city := query.Get("city")
	timestamp := query.Get("dt")

	if city == "" {
		_, err := w.Write([]byte(`{"error": "city field is required"}`))
		handleError(w, err)
		return
	}
	if timestamp == "" {
		_, err := w.Write([]byte(`{"error": "timestamp field is required"}`))
		handleError(w, err)
		return
	}

	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	dt, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		handleError(w, err)
		return
	}

	data, err := makeResponse(city, apiKey, dt.Unix())
	if err != nil {
		handleError(w, err)
		return
	}

	forecast := Forecast{city, "celsius", data}
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
