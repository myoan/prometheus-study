package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	counter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "my_counter",
		Help: "counter sample",
	})
	diceGaugeVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "dice",
		Help: "gauge sample",
	}, []string{"env", "side"})
)

func DiceGauge(env string, side int, ret float64) {
	diceGaugeVec.WithLabelValues(env, fmt.Sprintf("%d", side)).Set(ret)
}

func init() {
	prometheus.MustRegister(counter)
	prometheus.MustRegister(diceGaugeVec)
}

func diceRoll() {
	environments := []string{"dev", "stg", "prod"}
	sides := []int{6, 8, 12, 24}

	for _, env := range environments {
		for _, side := range sides {
			ret := rand.Intn(side)
			DiceGauge(env, side, float64(ret))
		}
	}
}

func main() {
	go func() {
		for {
			counter.Inc()
			diceRoll()
			time.Sleep(1 * time.Second)
		}
	}()

	fmt.Println("Start APP server")

	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Error: ", err)
	}
}
