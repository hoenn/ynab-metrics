package categories

import (
	"log"

	u "ynab-metrics/pkg/units"

	"github.com/prometheus/client_golang/prometheus"
	"go.bmvs.io/ynab"
	"go.bmvs.io/ynab/api/budget"
)

var categoryBudgeted = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "category_budget",
	Help: "Category budget gauge",
},
	[]string{"budget_name", "name"})
var categoryActivity = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "category_activity",
	Help: "Category activity gauge",
},
	[]string{"budget_name", "name"})

func init() {
	prometheus.MustRegister(categoryBudgeted)
	prometheus.MustRegister(categoryActivity)
}

//StartMetrics collects accounts metrics given a list of budgets
func StartMetrics(c ynab.ClientServicer, budgets []*budget.Budget) {
	log.Print("Getting Categories...")

	for _, b := range budgets {
		for _, c := range b.Categories {
			if !c.Deleted {
				categoryBudgeted.WithLabelValues(b.Name, c.Name).Set(float64(u.Dollars(c.Budgeted)))
				categoryActivity.WithLabelValues(b.Name, c.Name).Set(float64(u.Dollars(c.Activity)))
			}
		}
	}
}
