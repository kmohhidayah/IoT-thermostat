package broker

type Temperature struct {
	ID        string  `json:"id"`
	Temp      float64 `json:"temp"`
	Timestamp string  `json:"timestamp"`
}

type MaxTemperature struct {
	MaxTemp float64 `json:"maxtemp"`
}

type MinTemperature struct {
	MinTemp float64 `json:"mintemp"`
}

type LatestTemperature struct {
	LatestTemp float64 `json:"latesttemp"`
	LatestTime string  `json:"latesttime"`
}

//func
var MaxTemp []MaxTemperature
var Temps []Temperature
var Temp Temperature
