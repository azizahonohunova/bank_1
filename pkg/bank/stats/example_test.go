package stats

import (
	"fmt"

	"github.com/azizahonohunova/bank/pkg/bank/types"
)

func ExampleAvg() {
	payment := []types.Payment{
		{
			ID:       1000,
			Amount:   10_000_00,
			Category: "gold",
		},
		{
			ID:       2000,
			Amount:   20_000_00,
			Category: "gold",
		},
	}
	max := Avg(payment)
	fmt.Println(max)
	//Output:
	//1500000
}
