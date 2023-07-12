package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func serveMetrics(addr, path string) {
	const me = "serveMetrics"
	log.Printf("%s: starting metrics server at: %s %s", me, addr, path)
	http.Handle(path, promhttp.Handler())
	err := http.ListenAndServe(addr, nil)
	log.Fatalf("%s: ListenAndServe error: %v", me, err)
}

type metrics struct {
	latency *prometheus.HistogramVec
}

func newMetrics(namespace string, latencyBuckets []float64) *metrics {
	const me = "newMetrics"

	//
	// latency
	//

	latency := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "ping_requests_seconds",
			Help:      "Ping request duration in seconds.",
			Buckets:   latencyBuckets,
		},
		[]string{"outcome"},
	)

	if err := prometheus.Register(latency); err != nil {
		log.Fatalf("%s: latency was not registered: %s", me, err)
	}

	//
	// all metrics
	//

	m := &metrics{
		latency: latency,
	}

	return m
}

func (m *metrics) recordLatency(outcome string, latency time.Duration) {
	m.latency.WithLabelValues(outcome).Observe(float64(latency) / float64(time.Second))
}
