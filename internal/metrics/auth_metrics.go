package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type SessionCounter interface {
    CountActiveSessions() (int64, error)
}

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

func Register(counter SessionCounter) {
	activeSessions := prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "auth_active_sessions",
			Help: "Current number of active user sessions",
		},
		func() float64 {
			count, err := counter.CountActiveSessions()
			if err != nil {
				return 0
			}
			return float64(count)
		},
	)

	prometheus.MustRegister(
		LoginSuccessTotal,
		LoginFailureTotal,
		activeSessions,
		HTTPRequestsTotal,
		HTTPRequestDuration,
	)
}
