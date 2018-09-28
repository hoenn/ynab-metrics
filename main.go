package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go.bmvs.io/ynab"
	"go.bmvs.io/ynab/api/account"
)

var (
	addr  = flag.String("listen-address", ":8080", "The address to lsiten on for HTTP requests")
	token = flag.String("token", "", "Your YNAB access token")
)

var (
	xyzMetric = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "xyz_metric_gauge",
		Help: "XYZ gauge",
	})
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
	prometheus.MustRegister(xyzMetric)
	prometheus.MustRegister(rateLimitUsed)
	prometheus.MustRegister(rateLimitTotal)
}

func main() {
	flag.Parse()
	if *token == "" {
		panic("Must run with --token=abc... flag")
	}

	c := ynab.NewClient(*token)
	budgetSummaries, err := c.Budget().GetBudgets()
	if err != nil {
		panic(err)
	}

	var accts []*account.Account
	for _, b := range budgetSummaries {
		s, err := c.Budget().GetBudget(b.ID, nil)
		fmt.Println(err)
		accts = append(accts, s.Budget.Accounts...)
	}
	for _, a := range accts {
		fmt.Println(a)
	}

	rateLimitUsed.Set(float64(c.RateLimit().Used()))
	rateLimitTotal.Set(float64(c.RateLimit().Total()))

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
