package swedbank

// Transactions
type TransactionList struct {
	List []map[string]Transaction `json:"transactions"`
}

type Transaction struct {
	CreditDebit           string               `json:"credit_debit"`
	Amount                Amount               `json:"amount"`
	BookingDate           string               `json:"booking_date"`
	TransactionDate       string               `json:"transaction_date"`
	ValueDate             string               `json:"value_date"`
	RemittanceInformation string               `json:"remittance_information"`
	Balances              []map[string]Balance `json:"balances"`
}

// GetTransactions ...
// Fetch transactions from the account
func GetTransactions(acc *SingleAccount) (*TransactionList, error) {
	return nil, nil
}
