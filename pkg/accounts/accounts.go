package accounts

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"go.bmvs.io/ynab"
	"go.bmvs.io/ynab/api/budget"
	u "ynab-metrics/pkg/units"
)

var accountBalance = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "account_balance",
	Help: "Account balance gauge",
},
	[]string{"uuid", "name"})

func init() {
	prometheus.MustRegister(accountBalance)
}

//StartMetrics collects accounts metrics given a list of budgets
func StartMetrics(c ynab.ClientServicer, budgets []*budget.Budget) {
	log.Println("Getting Accounts...")

	for _, b := range budgets {
		for _, a := range b.Accounts {
			accountBalance.WithLabelValues(a.ID, a.Name).Set(float64(u.Dollars(a.Balance)))
		}
	}
}
