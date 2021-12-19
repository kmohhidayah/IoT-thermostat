package broker

type Temperature struct {
	ID        string  `json:"id"`
	Temp      float64 `json:"temp"`
	Timestamp string  `json:"timestamp"`
}

var Temps []Temperature
var Temp Temperature
