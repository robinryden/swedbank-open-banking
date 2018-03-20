package swedbank

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Transactions
type TransactionList struct {
	List map[string][]Transaction `json:"transactions"`
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
	payload, err := http.NewRequest("GET", accountAPI+acc.ID+"/transactions/?bic="+acc.BIC+"&date_from=2016-01-01&date_to=2018-01-01", nil)
	payload.Header.Add("Process-ID", acc.ProcessID)
	payload.Header.Add("Request-ID", acc.RequestID)
	payload.Header.Add("Authorization", acc.Authorization)

	req, err := client.Do(payload)

	if err != nil {
		fmt.Println("Error occured while trying to fetch from", accountAPI)
	}

	defer req.Body.Close()
	decoder := json.NewDecoder(req.Body)

	if err != nil {
		return nil, err
	}

	fmt.Println(decoder)

	transaction := TransactionList{}
	err = decoder.Decode(&transaction)

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
