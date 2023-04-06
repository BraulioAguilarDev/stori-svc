package report

type Balance struct {
	Month         string
	Debit, Credit int
	Subtotal      float64
}

type Average struct {
	Credit, Debit, NumberDebit, NumberCredit float64
}

type TransactionMonth struct {
	Month  string
	Number int
}

type Summary struct {
	Total, AverageDebitAmount, AverageCreditAmount float64
	Transactions                                   []TransactionMonth
}
