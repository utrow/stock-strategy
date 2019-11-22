package services

import (
	iconv "github.com/djimenez/iconv-go"
	"encoding/csv"
	"fmt"
	"os"
)

type CsvReader struct {
}

func (c *CsvReader) ReadCsv(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("Failed to close file:", filename)
		}
	}()

	converter, err := iconv.NewReader(file, "sjis", "utf-8")
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(converter)
	reader.FieldsPerRecord = -1
	return reader.ReadAll()
}
