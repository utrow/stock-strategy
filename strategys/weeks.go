package strategys

import (
	"fmt"
	"github.com/utrow/stock-strategy/models"
)

type Weeks struct {
	profitTotal    int
	purchasedPrice int
	totalCount     int
	winCount       int
	paymentMax     int
}

func (w *Weeks) TryStrategy1(h []models.History) {
	for i := 0; i < len(h)-1; i++ {
		if i < 10 {
			continue
		}

		today := h[i]
		tomorrow := h[i+1]

		_, currentWeekNo := today.Date.ISOWeek()
		_, tomorrowWeekNo := tomorrow.Date.ISOWeek()
		var previousWeekNo *int
		currentWeek := today
		var previousWeek models.History

		for j := 0; j < 10; j++ {
			cursorStock := h[i-j]
			_, cursorWeekNo := cursorStock.Date.ISOWeek()
			if cursorWeekNo == currentWeekNo {
				if currentWeek.Open > cursorStock.Open {
					currentWeek.Open = cursorStock.Open
				}
				if currentWeek.Low > cursorStock.Low {
					currentWeek.Low = cursorStock.Low
				}
				if currentWeek.High > cursorStock.High {
					currentWeek.High = cursorStock.High
				}
			} else {
				if previousWeekNo == nil {
					previousWeek = cursorStock
					previousWeekNo = &cursorWeekNo
				}

				if *previousWeekNo != cursorWeekNo {
					break
				}

				previousWeek.Date = cursorStock.Date
				if previousWeek.Open > cursorStock.Open {
					previousWeek.Open = cursorStock.Open
				}
				if previousWeek.Low > cursorStock.Low {
					previousWeek.Low = cursorStock.Low
				}
				if currentWeek.High < cursorStock.High {
					previousWeek.High = cursorStock.High
				}
			}
		}

		isChangeTomorrowWeek := currentWeekNo != tomorrowWeekNo

		if !w.isHold() && currentWeek.Open > currentWeek.Close && isChangeTomorrowWeek {
			w.buy(today.Close)
		} else if w.isHold() && isChangeTomorrowWeek {
			w.sellAll(today.Close)
		}
		w.holdPrice(today)
	}

	w.printResult()
}

func (w *Weeks) holdPrice(current models.History) {
	var profit int
	if w.purchasedPrice != 0 {
		profit = (current.Close - w.purchasedPrice) * 10
	}

	_, currentWeekNo := current.Date.ISOWeek()

	fmt.Println(
		current.Date.Format("2006-01-02"),
		current.Date.Weekday().String(), "\t(", currentWeekNo, ")\t",
		current.Close, "\t",
		"Hold: (", profit, ")\t",
		"Total (", w.profitTotal, ")",
	)
}

func (w *Weeks) buy(price int) {
	w.purchasedPrice = price
	w.totalCount += 1

	if w.paymentMax < price {
		w.paymentMax = price
	}
	fmt.Println("Buy:", price)
}

func (w *Weeks) sellAll(price int) {
	profit := (price - w.purchasedPrice) * 10
	if profit > 0 {
		w.winCount += 1
	}

	w.profitTotal += profit
	w.purchasedPrice = 0
	fmt.Println("Sell:", price, "(", profit, ")")
}

func (w *Weeks) isHold() bool {
	return w.purchasedPrice != 0
}

func (w *Weeks) printResult() {
	profitPer := float32(w.profitTotal) / float32(w.paymentMax*10)
	winningPer := float32(w.winCount) / float32(w.totalCount)
	fmt.Println("----------------------------------------")
	fmt.Println("Payment Total:", w.profitTotal, "/", w.paymentMax*10, "(", profitPer, ")")
	fmt.Println("Winning:", winningPer)
}
