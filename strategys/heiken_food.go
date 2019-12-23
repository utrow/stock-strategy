package strategys

import (
	"fmt"
	"github.com/utrow/stock-strategy/models"
	"strconv"
	"time"
)

type heikenFoodRow struct {
	Date  time.Time
	Open  int
	High  int
	Low   int
	Close int
}

type HeikenFood struct {
	heikenFoodHistories []heikenFoodRow
	profitTotal         int
	stockGettingPrice   int
	stockSize           int
}

func (r *HeikenFood) TryPrevChanged(h []models.History) {
	r.setHeikenFoodHistory(h)

	histories := r.heikenFoodHistories

	fmt.Println("Date\tClose\tTrend\tHold\tProfit")

	for i := 0; i < len(histories); i++ {
		if i < 4 {
			continue
		}

		current := h[i]
		prevHeiken := histories[i-2]
		prev2Heiken := histories[i-3]
		prev3Heiken := histories[i-4]

		isPrevUpTrend := prevHeiken.Open < prevHeiken.Close
		isPrev2UpTrend := prev2Heiken.Open < prev2Heiken.Close
		isPrev3UpTrend := prev3Heiken.Open < prev3Heiken.Close

		var displayTrend string
		if isPrevUpTrend {
			displayTrend = "🟩"
		} else {
			displayTrend = "🟥"
		}

		if r.stockSize == 0 && isPrevUpTrend && !isPrev2UpTrend && !isPrev3UpTrend {
			r.orderBuy(100, current.Open)
		} else if r.stockSize > 0 && !isPrevUpTrend && isPrev2UpTrend {
			r.orderSellAll(current.Open)
		}

		fmt.Println(
			current.Date.Format("2006-01-02"),
			current.Close,
			displayTrend,
			r.getStockProfit(current.Close),
			r.profitTotal,
		)
	}

	r.printResult()
}

func (r *HeikenFood) TryCurrentChanged(h []models.History) {
	r.setHeikenFoodHistory(h)

	histories := r.heikenFoodHistories

	fmt.Println("Date\tClose\tTrend\tHold\tProfit")

	for i := 0; i < len(histories); i++ {
		if i < 2 {
			continue
		}

		current := h[i]
		prevHeiken := histories[i-2]
		currentHeikenClose := (current.Open + current.High + current.Low + current.Close) / 4
		currentHeikenOpen := (prevHeiken.Open + prevHeiken.Close) / 2

		isPrevUpTrend := prevHeiken.Open < prevHeiken.Close
		isChangeUpTrend := currentHeikenOpen < currentHeikenClose

		var displayTrend string
		if isPrevUpTrend {
			displayTrend = "🟩"
		} else {
			displayTrend = "🟥"
		}

		if r.stockSize == 0 && isPrevUpTrend && isChangeUpTrend {
			r.orderBuy(100, current.Close)
		} else if r.stockSize > 0 && !isPrevUpTrend && !isChangeUpTrend {
			r.orderSellAll(current.Close)
		}

		fmt.Println(
			current.Date.Format("2006-01-02"),
			current.Close,
			displayTrend,
			r.getStockProfit(current.Close),
			r.profitTotal,
		)
	}

	r.printResult()
}

func (r *HeikenFood) setHeikenFoodHistory(h []models.History) {
	histories := make([]heikenFoodRow, 0)
	for i := 0; i < len(h); i++ {
		if i < 1 {
			continue
		}

		current := h[i]
		prev := h[i-1]

		if i == 1 {
			histories = append(histories, heikenFoodRow{
				Date:  current.Date,
				Open:  (prev.Open + prev.High + prev.Low + prev.Close) / 4,
				High:  current.High,
				Low:   current.Low,
				Close: (current.Open + current.High + current.Low + current.Close) / 4,
			})
		} else {
			prevHeikenFood := histories[len(histories)-1]
			histories = append(histories, heikenFoodRow{
				Date:  current.Date,
				Open:  (prevHeikenFood.Open + prevHeikenFood.Close) / 2,
				High:  current.High,
				Low:   current.Low,
				Close: (current.Open + current.High + current.Low + current.Close) / 4,
			})
		}
	}

	r.heikenFoodHistories = histories
}

func (r *HeikenFood) orderBuy(size int, price int) {
	fmt.Println("Buy:", price)
	r.stockGettingPrice = price
	r.stockSize = size
}

func (r *HeikenFood) orderSellAll(price int) {
	fmt.Println("Sell:", price, "("+strconv.Itoa(price-r.stockGettingPrice)+")")

	r.profitTotal += r.getStockProfit(price)

	r.stockGettingPrice = 0
	r.stockSize = 0
}

func (r *HeikenFood) getStockProfit(price int) int {
	return (price - r.stockGettingPrice) * r.stockSize
}

func (r *HeikenFood) printResult() {
	fmt.Println("----------------------------------------")
	fmt.Println("Profit:", 0, "(", 0, ")")
}
