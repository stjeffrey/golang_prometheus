package main

import (
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var requestsProcessed = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "test",
			Help: "test",
		},
		[]string{"status"},
)

func init() {
	prometheus.MustRegister(requestsProcessed)
}

func recordEvent() {
    requestsProcessed.WithLabelValues("test").Inc()
}

func eventHandler(w http.ResponseWriter, r *http.Request) {
    recordEvent()
    w.Write([]byte("Event recorded\n"))
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/record_event", eventHandler)
	http.ListenAndServe(":8080", nil)
}