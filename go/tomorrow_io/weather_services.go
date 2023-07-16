package tomorrow_io

import (
	"fmt"
	"net/http"
)

// GetWeatherForecast using the weather forecast API fetches up-to-date weather information for the specified location,
// including minute-by-minute forecasts for the next hour, hourly forecasts for the next 120 hours, and daily forecasts for the next 5 days.
//
// Docs: https://docs.tomorrow.io/reference/weather-forecast
func GetWeatherForecast() WeatherForecast {
	params :=
		map[string]any{
			"location": "arekere%20mico%20layout",
			"timesteps": []Timestep{
				Minutely, Hourly, Hourly,
			},
			"units": "metric",
		}
	url := encodeTomorrowUrl("weather/forecast", params)

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var weatherForecast WeatherForecast

	err = ParseJson(response.Body, &weatherForecast)

	if err != nil {
		fmt.Print("Error desirialization of WeatherForecast")
	}

	return weatherForecast
}

// GetRealtimeWeather using the Realtime Weather API fetches current weather
// information for the specified location in minute-by-minute temporal resolution.
//
// Docs: https://docs.tomorrow.io/reference/realtime-weather
func GetRealtimeWeather() RealtimeWeather {

	params :=
		map[string]any{
			"location": "arekere%20mico%20layout",
			"units":    "metric",
		}
	url := encodeTomorrowUrl("weather/realtime", params)

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var realtimeWeather RealtimeWeather

	err = ParseJson(response.Body, &realtimeWeather)

	if err != nil {
		fmt.Print("Error desirialization of RealtimeWeather")
	}

	return realtimeWeather
}

// GetWeatherHistory Using the recent history weather API fetches historical weather forecasts for the specified location,
// including hourly history for the last 24 hours, and daily history for the last day.
//
// Docs: https://docs.tomorrow.io/reference/weather-recent-history
func GetWeatherHistory() WeatherHistory {

	params :=
		map[string]any{
			"location": "arekere%20mico%20layout",
			"timesteps": []Timestep{
				Minutely, Hourly, Hourly,
			},
			"units": "metric",
		}

	url := encodeTomorrowUrl("weather/history/recent", params)

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var weatherHistory WeatherHistory

	err = ParseJson(response.Body, &weatherHistory)

	if err != nil {
		fmt.Print("Error desirialization of WeatherHistory")
	}

	return weatherHistory
}
