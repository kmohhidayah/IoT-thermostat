package http

import (
	"backend/broker"
	"encoding/json"
	"net/http"
)

func HttpServer() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		rw.Header().Add("Access-Control-Allow-Origin", "*")
		json.NewEncoder(rw).Encode(struct {
			Latest       broker.Temperature   `json:"latest"`
			Temperatures []broker.Temperature `json:"temperatures"`
		}{
			Latest:       broker.Temps[len(broker.Temps)-1],
			Temperatures: broker.Temps,
		})
	})
	http.ListenAndServe(":4000", nil)
}
