package accounts

import (
	"log"

	u "github.com/theoxifier/ynab-metrics/pkg/units"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/brunomvsouza/ynab.go"
	"github.com/brunomvsouza/ynab.go/api/budget"
)

var accountBalance = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "account_balance",
	Help: "Account balance gauge",
},
	[]string{"budget_name", "name"})

func init() {
	prometheus.MustRegister(accountBalance)
}

//StartMetrics collects accounts metrics given a list of budgets
func StartMetrics(c ynab.ClientServicer, budgets []*budget.Budget) {
	log.Println("Getting Accounts...")

	for _, b := range budgets {
		for _, a := range b.Accounts {
			accountBalance.WithLabelValues(b.Name, a.Name).Set(float64(u.Dollars(a.Balance)))
		}
	}
}
