package main

import (
	"fmt"
	"smart_vibes/tomorrow_io"
)

func main() {
	weatherForecast := tomorrow_io.GetWeatherForecast()
	fmt.Print(weatherForecast)

	realtimeWeather := tomorrow_io.GetRealtimeWeather()
	fmt.Print(realtimeWeather)

	weatherHistory := tomorrow_io.GetWeatherHistory()
	fmt.Print(weatherHistory)
}
