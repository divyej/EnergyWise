package services

import (
	"fmt"
	"net/http"
	"smart_vibes/models"
	"smart_vibes/utils"
)

// GetWeatherForecast fetches Weather Forecast Using the weather forecast API which gives access up-to-date weather information for location,
// including minute-by-minute forecasts for the next hour, hourly forecasts for the next 120 hours, and daily forecasts for the next 5 days.
//
// Docs: https://docs.tomorrow.io/reference/weather-forecast
// 
// Sample https://api.tomorrow.io/v4/weather/forecast?location=lake%20town&timesteps=1d&timesteps=1h&timesteps=1m&units=metric&apikey=<API_KEY>
func GetWeatherForecast() models.WeatherForecast {
	params :=
		map[string]any{
			"timesteps": []string{
				"1m", "1h", "1d",
			},
			"units":    "metric",
			"location": "arekere%20mico%20layout",
		}
	url := utils.EncodeParameters("weather/forecast?", params)

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var weatherForecast models.WeatherForecast

	err = utils.ParseJson(response.Body, &weatherForecast)

	if err != nil {
		fmt.Print("Error desirialization of WeatherForecast")
	}

	return weatherForecast
}
