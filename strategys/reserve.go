package strategys

import (
	"fmt"
	"github.com/utrow/stock-strategy/models"
)

type Reserve struct {
	paymentTotal int
	stockSize    int
}

func (r *Reserve) TryStrategy(h []models.History) {
	for i := 0; i < len(h); i++ {
		if i < 2 {
			continue
		}

		today := h[i]
		yesterday := h[i-1]

		if yesterday.Low > today.Open {
			r.order(1, today.Low)
			fmt.Println("Order", today.Date)
		}

		stockTotal := r.stockSize * today.Close
		fmt.Println("Profit:", stockTotal-r.paymentTotal)
	}

	r.printResult(h[len(h)-1].Close)
}

func (r *Reserve) order(size int, price int) {
	r.paymentTotal += size * price
	r.stockSize += size
}

func (r *Reserve) printResult(latestPrice int) {
	stockTotal := r.stockSize * latestPrice
	profit := stockTotal - r.paymentTotal
	profitPer := float32(profit) / float32(r.paymentTotal)
	fmt.Println("Payment Total:", r.paymentTotal)
	fmt.Println("Stock Total:", stockTotal)
	fmt.Println("----------------------------------------")
	fmt.Println("Profit:", profitPer, "(", profit, ")")
}
