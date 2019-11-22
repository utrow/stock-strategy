package main

import (
	"fmt"
	"github.com/utrow/stock-storategy/interfaces"
	"github.com/utrow/stock-storategy/models"
	"github.com/utrow/stock-storategy/services"
	"github.com/utrow/stock-storategy/strategys"
)

func main() {
	service := services.MakeService()
	strategy := strategys.MakeStrategy()

	input := interfaces.GetArgInputs()
	if input == nil {
		fmt.Println("Error: Invalid input args.")
		return
	}

	i := *input
	rawHistories, err := service.CsvReader.ReadCsv(i.HistoryFile)
	if err != nil {
		fmt.Println("Failed ReadCsv.", err)
		return
	}

	histories := models.HistoryFactory(rawHistories)

	strategy.Reserve.TryStrategy(histories)
}
