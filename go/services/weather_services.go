package services

import (
	"fmt"
	"net/http"
	"smart_vibes/models"
	"smart_vibes/utils"
)

// GetWeatherForecast using the weather forecast API fetches up-to-date weather information for location,
// including minute-by-minute forecasts for the next hour, hourly forecasts for the next 120 hours, and daily forecasts for the next 5 days.
//
// Docs: https://docs.tomorrow.io/reference/weather-forecast
func GetWeatherForecast() models.WeatherForecast {
	params :=
		map[string]any{
			"location": "arekere%20mico%20layout",
			"timesteps": []string{
				"1m", "1h", "1d",
			},
			"units": "metric",
		}
	url := utils.EncodeTomorrowUrl("weather/forecast", params)

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

// GetRealtimeWeather using the Realtime Weather API fetches current weather
// information for the specified location in minute-by-minute temporal resolution.
//
// Docs: https://docs.tomorrow.io/reference/realtime-weather
func GetRealtimeWeather() models.RealtimeWeather {

	params :=
		map[string]any{
			"location": "arekere%20mico%20layout",
			"units":    "metric",
		}
	url := utils.EncodeTomorrowUrl("weather/realtime", params)

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var realtimeWeather models.RealtimeWeather

	err = utils.ParseJson(response.Body, &realtimeWeather)

	if err != nil {
		fmt.Print("Error desirialization of WeatherForecast")
	}

	return realtimeWeather
}
