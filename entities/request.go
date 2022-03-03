package entities

import (
	"encoding/json"
	"log"
	"net/http"
)

func (f Forcast) GetForecast()  {
	response, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=japan&appid=f2ab9d6dd296f998058bd6b4e3d0deb4")
	if err != nil {
		log.Println("can't reach server, ", err.Error())
		return
	}
	w:= Forcast{}
	err = json.NewDecoder(response.Body).Decode(&w)
	if err != nil {
		log.Println(err.Error())
		return
	}
}