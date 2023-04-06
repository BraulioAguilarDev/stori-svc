package reader

import (
	"encoding/csv"
	"io"
	"os"
)

type Data struct {
	ID        string
	Date      string
	Amount    string
	AccountID string
	BankName  string
	Number    string
	Currency  string
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

		// TODO: cast from string to int,data,float
		txns = append(txns, Data{
			ID:        row[0],
			Date:      row[1],
			Amount:    row[2],
			AccountID: row[3],
			BankName:  row[4],
			Number:    row[5],
			Currency:  row[6],
		})
	}

	return txns, nil
}
