package reader

import (
	"encoding/csv"
	"io"
	"os"
)

type Data struct {
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
			Date:         row[0],
			Amount:       row[1],
			BankName:     row[2],
			Number:       row[3],
			Currency:     row[4],
			AccountName:  row[5],
			AccountEmail: row[6],
		})
	}

	return txns, nil
}
