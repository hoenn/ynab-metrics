package transactions

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"go.bmvs.io/ynab"
	"go.bmvs.io/ynab/api/budget"
	"go.bmvs.io/ynab/api/transaction"
	"log"
	u "ynab-metrics/pkg/units"
)

var budgetTransactions = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "budget_transactions",
	Help: "Budget transactions",
},
	[]string{"budget_uuid", "payee_name", "category_name", "account_name", "transaction_id"})

func init() {
	prometheus.MustRegister(budgetTransactions)
}

type transactionFromBudget struct {
	transaction *transaction.Transaction
	budgetID    string
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
		fmt.Println(t)
		budgetTransactions.WithLabelValues(
			t.budgetID,
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
			budgetID:    b.ID,
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
