package stats

import "github.com/s-zer0/bank/v2/pkg/types"

const statusFails = types.StatusFail

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	avg := types.Money(0)
	quantity := types.Money(0)
	for _, v := range payments {
		if v.Status != statusFails {
		avg += v.Amount
		quantity ++
		}
	}
	return avg/quantity
}

// TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	totalCategory := types.Money(0)
	
	for _, payment := range payments {
		if payment.Status != statusFails {
			if payment.Category == category {
				totalCategory += payment.Amount
		}
	}
	}
	return totalCategory
}