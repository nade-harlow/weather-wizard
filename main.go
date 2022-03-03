package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/nade-harlow/weather-wizard/entities"
)

func main() {
	_ = godotenv.Load()
	f := entities.Forecast{}
	var city string
	fmt.Print("Enter a city name: ")
	fmt.Scan(&city)
	f.GetForecast(city)
}
