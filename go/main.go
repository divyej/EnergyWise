package main

import (
	"fmt"
	"smart_vibes/services"
)

func main() {
	weatherForecast := services.GetWeatherForecast()
	fmt.Print(weatherForecast)

	realtimeWeather := services.GetRealtimeWeather()
	fmt.Print(realtimeWeather)

	weatherHistory := services.GetWeatherHistory()
	fmt.Print(weatherHistory)
}
