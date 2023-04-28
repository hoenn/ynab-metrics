package budgets

import (
	"fmt"
	"log"
	"github.com/brunomvsouza/ynab.go"
	"github.com/brunomvsouza/ynab.go/api/budget"
)

//GetBudgets returns the list of budgets associated with user token
func GetBudgets(c ynab.ClientServicer) []*budget.Budget {
	log.Print("Getting Budgets...")
	budgetSummaries, err := c.Budget().GetBudgets()
	if err != nil {
		log.Print(err)
		log.Fatal("Unable to get budgets, check your access token")
	}

	var budgets []*budget.Budget
	for _, b := range budgetSummaries {
		b, err := c.Budget().GetBudget(b.ID, nil)
		if err != nil {
			log.Print(fmt.Sprintf("Unable to get data for budget: %s", b.Budget.ID))
		}
		budgets = append(budgets, b.Budget)
	}
	return budgets
}
