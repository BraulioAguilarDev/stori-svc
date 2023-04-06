package reader

import (
	"encoding/csv"
	"io"
	"os"
)

type Data struct {
	ID           string
	Date         string
	Amount       string
	BankName     string
	Number       string
	Currency     string
	AccountName  string
	AccountEmail string
}

func ReadFile(filename string) ([]Data, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var txns []Data
	csvReader := csv.NewReader(file)
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		txns = append(txns, Data{
			ID:           row[0],
			Date:         row[1],
			Amount:       row[2],
			BankName:     row[3],
			Number:       row[4],
			Currency:     row[5],
			AccountName:  row[6],
			AccountEmail: row[7],
		})
	}

	return txns, nil
}
