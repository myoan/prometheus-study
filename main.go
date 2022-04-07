package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	counter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myCounter",
		Help: "counter sample",
	})
)

func main() {
	go func() {
		for {
			counter.Inc()
			time.Sleep(1 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8000", nil)
}
