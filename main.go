package main

import (
	"flag"
	"log"
	"net/http"
	"time"
	"ynab-metrics/pkg/accounts"
	"ynab-metrics/pkg/budgets"
	"ynab-metrics/pkg/ratelimit"
	"ynab-metrics/pkg/transactions"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go.bmvs.io/ynab"
)

var (
	addr  = flag.String("listen-address", ":8080", "The address to lsiten on for HTTP requests")
	token = flag.String("token", "", "Your YNAB access token")
)

func main() {
	flag.Parse()
	if *token == "" {
		log.Fatal("Must run with --token=abc... flag")
	}

	c := ynab.NewClient(*token)

	go func() {
		for {
			budgets := budgets.GetBudgets(c)
			accounts.StartMetrics(c, budgets)
			transactions.StartMetrics(c, budgets)
			ratelimit.StartMetrics(c)

			time.Sleep(time.Duration(90 * time.Second))
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
