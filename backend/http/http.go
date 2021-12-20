package http

import (
	"backend/broker"
	"encoding/json"
	"net/http"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		rw.Header().Add("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(rw, r)
	})
}
func HttpServer() {
	router := http.NewServeMux()
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode(struct {
			Latest       broker.Temperature   `json:"latest"`
			Temperatures []broker.Temperature `json:"temperatures"`
		}{
			Latest:       broker.Temps[len(broker.Temps)-1],
			Temperatures: broker.Temps,
		})
	})
	router.HandleFunc("/latest", func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode(broker.LatestTemperature{
			LatestTemp: broker.Temp.Temp,
			LatestTime: broker.Temp.Timestamp,
		})
	})

	router.HandleFunc("/max", func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode(broker.MaxTemperature{MaxTemp: FindMax()})
	})
	router.HandleFunc("/min", func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode(broker.MinTemperature{MinTemp: FindMin()})
	})
	routerWithMiddleware := CORSMiddleware(router)
	http.ListenAndServe(":9090", routerWithMiddleware)
}
