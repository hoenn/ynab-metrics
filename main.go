package main

import (
	"flag"
	"log"
	"net/http"
	"time"
	"ynab-metrics/pkg/accounts"
	"ynab-metrics/pkg/budgets"
	"ynab-metrics/pkg/categories"
	"ynab-metrics/pkg/config"
	"ynab-metrics/pkg/ratelimit"
	"ynab-metrics/pkg/transactions"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go.bmvs.io/ynab"
)

var cfgFile = flag.String("config", "config.json", "The configuration file to use for exporting")

func main() {
	flag.Parse()
	cfg := config.ParseConfig(*cfgFile)
	if cfg.AccessToken == "" {
		log.Fatal("User token is empty")
	}

	c := ynab.NewClient(cfg.AccessToken)

	go func() {
		for {
			budgets := budgets.GetBudgets(c)
			accounts.StartMetrics(c, budgets)
			categories.StartMetrics(c, budgets)
			if cfg.GetTrans {
				transactions.StartMetrics(c, budgets)
			}
			ratelimit.StartMetrics(c)

			time.Sleep(time.Duration(cfg.IntervalSeconds) * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
