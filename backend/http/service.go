package http

import (
	"backend/broker"
)

func FindMax() float64 {
	var Value []float64
	var largerNumber, temp float64
	for _, n := range broker.Temps {
		Value = append(Value, n.Temp)
		if n.Temp > temp {
			temp = n.Temp
			largerNumber = temp
		}

	}
	return largerNumber
}

func FindMin() float64 {
	var Value []float64
	var smallerNumber, temp float64
	for _, n := range broker.Temps {
		Value = append(Value, n.Temp)
		if n.Temp < temp {
			temp = n.Temp
			smallerNumber = temp
		}

	}
	return smallerNumber
}
