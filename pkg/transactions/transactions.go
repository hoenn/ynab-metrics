package transactions

import (
	"fmt"
	"log"

	u "github.com/theoxifier/ynab-metrics/pkg/units"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/brunomvsouza/ynab.go"
	"github.com/brunomvsouza/ynab.go/api/budget"
	"github.com/brunomvsouza/ynab.go/api/transaction"
)

var budgetTransactions = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "budget_transactions",
	Help: "Budget transactions",
},
	[]string{"budget_name", "payee_name", "category_name", "account_name", "transaction_id"})

func init() {
	prometheus.MustRegister(budgetTransactions)
}

type transactionFromBudget struct {
	transaction *transaction.Transaction
	budgetName  string
}

//StartMetrics abcd...
func StartMetrics(c ynab.ClientServicer, budgets []*budget.Budget) {
	log.Print("Getting Transactions...")

	var transactions []*transactionFromBudget
	for _, b := range budgets {
		ts, err := c.Transaction().GetTransactions(b.ID, nil)
		if err != nil {
			log.Print(fmt.Sprintf("Unable to get transactions for budget: %s", b.ID))
		}

		tfbs := addBudgetIDs(ts, b)

		transactions = append(transactions, tfbs...)
	}

	for _, t := range transactions {
		budgetTransactions.WithLabelValues(
			t.budgetName,
			safe(t.transaction.PayeeName),
			safe(t.transaction.CategoryName),
			t.transaction.AccountName,
			t.transaction.ID).Set(float64(u.Dollars(t.transaction.Amount)))
	}
}

func addBudgetIDs(ts []*transaction.Transaction, b *budget.Budget) []*transactionFromBudget {
	var tfbs []*transactionFromBudget
	for _, t := range ts {
		tfbs = append(tfbs, &transactionFromBudget{
			transaction: t,
			budgetName:  b.Name,
		})
	}
	return tfbs
}

func safe(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
