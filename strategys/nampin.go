package strategys

import "github.com/utrow/stock-strategy/models"

type Nampin struct {
	profitTotal    int
	purchasedPrice int
	totalCount     int
	winCount       int
	paymentTotal   int
}

func (n *Nampin) TryStrategy1(h []models.History) {

}
