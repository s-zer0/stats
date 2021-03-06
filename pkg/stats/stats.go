package stats

import "github.com/s-zer0/bank/v2/pkg/types"

// const statusFails = types.StatusFail

// // Avg рассчитывает среднюю сумму платежа.
// func Avg(payments []types.Payment) types.Money {
// 	avg := types.Money(0)
// 	quantity := types.Money(0)
// 	for _, v := range payments {
// 		if v.Status != statusFails {
// 		avg += v.Amount
// 		quantity ++
// 		}
// 	}
// 	return avg/quantity
// }

// // TotalInCategory находит сумму покупок в определённой категории.
// func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
// 	totalCategory := types.Money(0)
	
// 	for _, payment := range payments {
// 		if payment.Status != statusFails {
// 			if payment.Category == category {
// 				totalCategory += payment.Amount
// 		}
// 	}
// 	}
// 	return totalCategory
// }

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	quantity := map[types.Category]types.Money{}	
	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
		quantity[payment.Category] += 1
	}
	for key := range categories {
		categories[key] = (categories[key]/quantity[key])
	}
	return categories
}

func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	categoriesPeriod := map[types.Category]types.Money{}
	for cat := range first{
		_, v := second[cat]
		if !v {
			categoriesPeriod[cat]=first[cat]*-1
			continue	
		}
		for cat2 := range second{
			_, v2 := first[cat2]
			if !v2 {
				categoriesPeriod[cat2]=second[cat2]
				continue	
			}
		}
		categoriesPeriod[cat]=second[cat]-first[cat]
	}
	return categoriesPeriod
}