package report

type Balance struct {
	Month         string
	Debit, Credit int
	Subtotal      float64
}

type Average struct {
	Credit, Debit, NumberDebit, NumberCredit float64
}

type AccountSummary struct {
	Name, Email string
	Balances    []Balance
	Average     Average
}

type TransactionMonth struct {
	Month  string
	Number int
}

type Summary struct {
	Name                                           string
	Email                                          string
	Total, AverageDebitAmount, AverageCreditAmount float64
	Transactions                                   []TransactionMonth
}
