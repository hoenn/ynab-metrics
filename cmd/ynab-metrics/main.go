package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/theoxifier/ynab-metrics/pkg/accounts"
	"github.com/theoxifier/ynab-metrics/pkg/budgets"
	"github.com/theoxifier/ynab-metrics/pkg/categories"
	"github.com/theoxifier/ynab-metrics/pkg/config"
	"github.com/theoxifier/ynab-metrics/pkg/ratelimit"
	"github.com/theoxifier/ynab-metrics/pkg/transactions"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/brunomvsouza/ynab.go"
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
