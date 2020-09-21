package stats

import "github.com/azizahonohunova/bank/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {

	count := len(payments)
	sum := 0
	for i := 0; i < count; i++ {
		sum += int(payments[i].Amount)
	}
	avg := sum / count
	return types.Money(avg)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {

	sum := 0
	for _, card := range payments {

		if card.Category == category {
			sum = sum + int(card.Amount)
		}
	}
	return types.Money(sum)
}
