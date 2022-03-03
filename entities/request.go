package entities

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func (f *Forecast) GetForecast(city string) {
	response, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, os.Getenv("APIKEY")))
	if err != nil {
		log.Println("can't reach server, ", err.Error())
		return
	}
	w := Forecast{}
	err = json.NewDecoder(response.Body).Decode(&w)
	if err != nil {
		log.Println(err.Error())
		return
	}
	f.New(w)
}

func (f *Forecast) New(forecast Forecast) {
	f.ID = forecast.ID
	f.Weather = forecast.Weather
	f.CityName = forecast.CityName
	f.Temperature = forecast.Temperature
	f.Sys = forecast.Sys
	f.Timezone = forecast.Timezone
}

func (f Forecast) Print() {
	if f.CityName == "" {
		log.Println("Error: Enter a city name")
		return
	}
	fmt.Printf("%s %s's Weather Forecast %s\n", strings.Repeat("-", 20), f.CityName, strings.Repeat("-", 20))
	fmt.Printf("City: %s\n", f.CityName)
	fmt.Printf("Country: %v\n", f.Sys.Country)
	fmt.Printf("Weather: %v (%s)\n", f.Weather[0].Main, f.Weather[0].Description)
	fmt.Printf("Temperature: %vF\n", f.Temperature.Temp)
	fmt.Printf("Timezone: %v\n", f.Timezone)
}
