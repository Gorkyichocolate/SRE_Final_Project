package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (

	// =========================
	// HTTP REQUESTS
	// =========================

	HttpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)

	// =========================
	// REQUEST DURATION
	// =========================

	HttpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request latency",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	// =========================
	// ORDERS
	// =========================

	OrdersCreatedTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "orders_created_total",
			Help: "Total created orders",
		},
	)

	// =========================
	// AUTH
	// =========================

	AuthRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "auth_requests_total",
			Help: "Authentication requests",
		},
		[]string{"type"},
	)

	// =========================
	// HEAVY ENDPOINT
	// =========================

	HeavyRequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "heavy_requests_total",
			Help: "Heavy endpoint requests",
		},
	)
)

func InitMetrics() {

	prometheus.MustRegister(
		HttpRequestsTotal,
		HttpRequestDuration,
		OrdersCreatedTotal,
		AuthRequestsTotal,
		HeavyRequestsTotal,
	)
}
