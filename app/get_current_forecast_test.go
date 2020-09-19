package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCurrentForecast(t *testing.T) {
	server := httptest.NewServer(Handlers())
	defer server.Close()

	city := "Omsk"
	res, err := http.Get(fmt.Sprintf("%s/v1/forecast/?city=%s", server.URL, city))

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Status is not OK")
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Fatal(err)
	}

	contentType := res.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Fatalf("Expected json response, got: %s", contentType)
	}

	expectedUnit := "celsius"
	expectedTemperature := 4.4
	forecast := Forecast{}
	err = json.Unmarshal(body, &forecast)

	if err != nil {
		t.Fatal(err)
	}

	if forecast.City != city {
		t.Fatalf("Expected %s, got %s", city, forecast.City)
	}
	if forecast.Unit != expectedUnit {
		t.Fatalf("Expected %s, got %s", expectedUnit, forecast.Unit)
	}
	if forecast.Temperature != expectedTemperature {
		t.Fatalf("Expected %d, got %d", expectedTemperature, forecast.Temperature)
	}
}
