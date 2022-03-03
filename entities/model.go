package entities

type Forecast struct {
	ID          int32       `json:"id"`
	CityName    string      `json:"name"`
	Weather     []Weather   `json:"weather"`
	Temperature Temperature `json:"main"`
	Sys         System      `json:"sys"`
	Timezone    int32       `json:"timezone"`
}

type Weather struct {
	Id          int16  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Temperature struct {
	Temp     float32 `json:"temp"`
	TempMin  float32 `json:"temp_min"`
	TempMax  float32 `json:"temp_max"`
	Pressure float32 `json:"pressure"`
}

type System struct {
	Country string `json:"country"`
}
