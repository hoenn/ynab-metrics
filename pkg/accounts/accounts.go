package accounts

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"go.bmvs.io/ynab"
	"go.bmvs.io/ynab/api/budget"
)

var accountBalance = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "account_balance",
	Help: "Account balance gauge",
},
	[]string{"uuid"})

func init() {
	prometheus.MustRegister(accountBalance)
}

//StartMetrics collects accounts metrics given a list of budgets
func StartMetrics(c ynab.ClientServicer, budgets []*budget.Budget) {
	fmt.Println("Accounts...")

	for _, b := range budgets {
		for _, a := range b.Accounts {
			accountBalance.WithLabelValues(a.ID).Set(float64(dollars(a.Balance)))
		}
	}
}
func dollars(milliunits int64) int64 {
	return milliunits / 1000
}
