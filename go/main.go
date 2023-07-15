package main

import (
	"fmt"
	"smart_vibes/services"
)

func main() {
	weatherForecast := services.GetWeatherForecast()
	fmt.Print(weatherForecast)
}
