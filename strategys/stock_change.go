package strategys

import (
	"fmt"
	"github.com/utrow/stock-strategy/models"
	"math"
)

type StockChange struct {
	profitTotal    int
	purchasedPrice int
	totalCount     int
	winCount       int
	paymentTotal   int
}

func (s *StockChange) TryStrategy1(h []models.History) {
	for i := 0; i < len(h); i++ {
		if i < 1 {
			continue
		}

		today := h[i]
		yesterday := h[i-1]

		todayYoSen := today.Open < today.Close
		rangeUp := math.Abs(float64(today.Close)-float64(today.Open)) > math.Abs(float64(yesterday.Close)-float64(yesterday.Open))

		if !s.isHold() && todayYoSen && rangeUp {
			s.buy(today.Close)
		} else if s.isHold() {
			s.sellAll(today.Close)
		}
		s.holdPrice(today)
	}

	s.printResult()
}

func (s *StockChange) TryStrategy2(h []models.History) {
	for i := 0; i < len(h); i++ {
		if i < 1 {
			continue
		}

		today := h[i]
		yesterday := h[i-1]

		yesterdayYoSen := yesterday.Open < yesterday.Close

		if !s.isHold() && yesterdayYoSen {
			s.buy(today.Open)
		} else if s.isHold() {
			s.sellAll(today.Open)
		}
		s.holdPrice(today)
	}

	s.printResult()
}

func (s *StockChange) TryStrategy3(h []models.History) {
	for i := 0; i < len(h); i++ {
		if i < 5 {
			continue
		}

		today := h[i]

		weekUp := h[i-5].Open < today.Close
		todayYoSen := today.Open < today.Close

		if !s.isHold() && todayYoSen && weekUp {
			s.buy(today.Close)
		} else if s.isHold() && !weekUp {
			s.sellAll(today.Close)
		}
		s.holdPrice(today)
	}

	s.printResult()
}

func (s *StockChange) holdPrice(current models.History) {
	profit := (current.Close - s.purchasedPrice) * 10

	fmt.Println(
		current.Date.Format("2006-01-02"),
		current.Close,
		profit,
		s.profitTotal,
	)
}

func (s *StockChange) buy(price int) {
	s.purchasedPrice = price
	s.paymentTotal += price
	s.totalCount += 1
	fmt.Println("Buy:", price)
}

func (s *StockChange) sellAll(price int) {
	profit := (price - s.purchasedPrice) * 10
	if profit > 0 {
		s.winCount += 1
	}

	s.profitTotal += profit
	s.purchasedPrice = 0
	fmt.Println("Sell:", price, "(", profit, ")")
}

func (s *StockChange) isHold() bool {
	return s.purchasedPrice != 0
}

func (s *StockChange) printResult() {
	profitPer := float32(s.profitTotal) / float32(s.paymentTotal)
	winningPer := float32(s.winCount) / float32(s.totalCount)
	fmt.Println("----------------------------------------")
	fmt.Println("Payment Total:", s.profitTotal, "/", s.paymentTotal, "(", profitPer, ")")
	fmt.Println("Winning:", winningPer)
}
