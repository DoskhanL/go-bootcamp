package weather

import (
	"encoding/json"
	"net/http"
)

// Weather struct for weather forecast
type Weather struct {
	City     string `json:"city"`
	Forecast string `json:"forecast"`
}

// GetWeather function getting weather from internet
func GetWeather(url string) (*Weather, error) {

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var weather Weather
	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		return nil, err
	}

	return &weather, nil
}
