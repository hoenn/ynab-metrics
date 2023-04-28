package ratelimit

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/brunomvsouza/ynab.go"
)

var (
	rateLimitUsed = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "rate_limit_used",
		Help: "Rate limit used of YNAB API",
	})
	rateLimitTotal = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "rate_limit_total",
		Help: "Rate limit total of YNAB API",
	})
)

func init() {
	prometheus.MustRegister(rateLimitUsed)
	prometheus.MustRegister(rateLimitTotal)
}

//StartMetrics writes rate limiting metrics from the client response
func StartMetrics(c ynab.ClientServicer) {
	log.Print("Getting Rate limting metrics...")
	rateLimitUsed.Set(float64(c.RateLimit().Used()))
	rateLimitTotal.Set(float64(c.RateLimit().Total()))
}
