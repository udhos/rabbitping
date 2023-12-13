package main

import (
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

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
