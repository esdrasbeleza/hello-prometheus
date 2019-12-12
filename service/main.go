package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			value := float64(time.Now().Unix()%20) * 10.0
			variatingNumber.Set(value)
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})

	variatingNumber = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "myapp_variating_number",
		Help: "Some variating stats",
	})
)

func main() {
	prometheus.MustRegister(opsProcessed)
	prometheus.MustRegister(variatingNumber)
	prometheus.MustRegister(prometheus.NewBuildInfoCollector())
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
