package main

import (
	"time"

	"github.com/udhos/boilerplate/envconfig"
)

type config struct {
	amqpURL               string
	interval              time.Duration
	timeout               time.Duration
	failureThreshold      int
	restartDeploy         string
	restartNamespace      string
	metricsAddr           string
	metricsPath           string
	metricsNamespace      string
	metricsLatencyBuckets []float64
	healthAddr            string
	healthPath            string
}

func getConfig(roleSessionName string) config {

	env := envconfig.NewSimple(roleSessionName)

	return config{
		amqpURL:               env.String("AMQP_URL", "amqp://guest:guest@rabbitmq:5672/"),
		interval:              env.Duration("INTERVAL", 10*time.Second),
		timeout:               env.Duration("TIMEOUT", 5*time.Second),
		failureThreshold:      env.Int("FAILURE_THRESHOLD", 6),
		restartDeploy:         env.String("RESTART_DEPLOY", ""),
		restartNamespace:      env.String("RESTART_NAMESPACE", "default"),
		metricsAddr:           env.String("METRICS_ADDR", ":3000"),
		metricsPath:           env.String("METRICS_PATH", "/metrics"),
		metricsNamespace:      env.String("METRICS_NAMESPACE", ""),
		metricsLatencyBuckets: env.Float64Slice("METRICS_BUCKETS_LATENCY", []float64{0.0005, 0.001, 0.0025, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, .5, 1, 2.5, 5}),
		healthAddr:            env.String("HEALTH_ADDR", ":8888"),
		healthPath:            env.String("HEALTH_PATH", "/health"),
	}
}
