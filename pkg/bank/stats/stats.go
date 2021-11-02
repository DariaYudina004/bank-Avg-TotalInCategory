package stats

import "github.com/DariaYudina004/bank/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	var sum types.Money

	for _, payment := range payments {
		if payment.Amount>0 {
			sum+=payment.Amount 
	}
}
	sum = sum /types.Money(len(payments)) 
	return sum
}