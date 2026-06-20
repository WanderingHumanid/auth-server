package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	LoginSuccessTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "auth_login_success_total",
			Help: "Total number of successful user logins",
		})

	LoginFailureTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "auth_login_failure_total",
			Help: "Total number of failed user login attempts",
		})
	ActiveSessions = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "auth_active_sessions",
			Help: "Current number of active user sessions",
		},
	)
	HTTPRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "auth_http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)
	HTTPRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request latency",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)
)

func Register() {
	prometheus.MustRegister(
		LoginSuccessTotal,
		LoginFailureTotal,
		ActiveSessions,
		HTTPRequestsTotal,
		HTTPRequestDuration,
	)
}
