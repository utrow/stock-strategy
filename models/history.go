package models

import (
	"fmt"
	"strconv"
	"time"
)

type History struct {
	Date  time.Time
	Open  int
	High  int
	Low   int
	Close int
}

func HistoryFactory(csvStringRows [][]string) (histories []History) {
	for i, row := range csvStringRows {

		headerCount := 2
		if i < headerCount {
			continue
		}

		if len(row) < 7 {
			fmt.Println("Missing field count on line", i+1)
			continue
		}

		date, err := time.Parse("2006-01-02", row[0])
		if err != nil {
			fmt.Println("Missing date on line", i+1, err)
			continue
		}

		openFloat, err := strconv.ParseFloat(row[1], 32)
		if err != nil {
			fmt.Println("Missing open on line", i+1, err)
			continue
		}

		highFloat, err := strconv.ParseFloat(row[2], 32)
		if err != nil {
			fmt.Println("Missing high on line", i+1, err)
			continue
		}

		lowFloat, err := strconv.ParseFloat(row[3], 32)
		if err != nil {
			fmt.Println("Missing low on line", i+1, err)
			continue
		}

		closeFloat, err := strconv.ParseFloat(row[4], 32)
		if err != nil {
			fmt.Println("Missing close on line", i+1, err)
			continue
		}

		histories = append(histories, History{
			Date: date,
			Open: int(openFloat * 10),
			High: int(highFloat * 10),
			Low: int(lowFloat * 10),
			Close: int(closeFloat * 10),
		})
	}
	return histories
}
