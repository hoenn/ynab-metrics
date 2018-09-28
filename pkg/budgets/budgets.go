package budgets

import (
	"fmt"
	"go.bmvs.io/ynab"
	"go.bmvs.io/ynab/api/budget"
	"log"
)

//GetBudgets returns the list of budgets associated with user token
func GetBudgets(c ynab.ClientServicer) []*budget.Budget {
	fmt.Println("Getting budgets...")
	budgetSummaries, err := c.Budget().GetBudgets()
	if err != nil {
		log.Fatal("Unable to get budgets, check your access token")
	}

	var budgets []*budget.Budget
	for _, b := range budgetSummaries {
		b, err := c.Budget().GetBudget(b.ID, nil)
		if err != nil {
			log.Print(fmt.Sprintf("Unable to data for budget: %s", b.Budget.ID))
		}
		budgets = append(budgets, b.Budget)
	}
	return budgets
}
