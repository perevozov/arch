package service

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	opsProcessed = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "myapp_processed_ops_total",
			Help: "The total number of processed events",
		},
		[]string{"method", "endpoint", "http_response"},
	)

	requestLatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "app_request_latency_seconds",
			Help: "Application request latency",
		},
		[]string{"method", "endpoint"},
	)
)
