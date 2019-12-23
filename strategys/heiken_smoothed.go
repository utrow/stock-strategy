package strategys

import (
	"fmt"
	"github.com/utrow/stock-strategy/models"
)

type heikenSmoothedRow struct {
}

type HeikenSmoothed struct {
	heikenSmoothedHistories []heikenSmoothedRow
	paymentTotal            int
	stockSize               int
}

func (r *HeikenSmoothed) TryStrategy(h []models.History) {
	r.setHeikenSmoothedHistory(h)
}

func (r *HeikenSmoothed) setHeikenSmoothedHistory(h []models.History) {
	var histories []heikenSmoothedRow
	for i := 0; i < len(h); i++ {
	}
	r.heikenSmoothedHistories = histories
}

func (r *HeikenSmoothed) order(size int, price int) {
	r.paymentTotal += size * price
	r.stockSize += size
}

func (r *HeikenSmoothed) printResult(latestPrice int) {
	stockTotal := r.stockSize * latestPrice
	profit := stockTotal - r.paymentTotal
	profitPer := float32(profit) / float32(r.paymentTotal)
	fmt.Println("Payment Total:", r.paymentTotal)
	fmt.Println("Stock Total:", stockTotal)
	fmt.Println("----------------------------------------")
	fmt.Println("Profit:", profitPer, "(", profit, ")")
}
