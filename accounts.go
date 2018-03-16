package swedbank

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	accounts = "https://psd2.api.swedbank.com/sandbox/v1/accounts/?bic=SANDSESS"
	account  = "https://psd2.api.swedbank.com/sandbox/v1/accounts/"
)

var (
	client = &http.Client{}
)

// Accounts
type SingleAccount struct {
	ID            string
	BIC           string
	RequestID     string
	ProcessID     string
	Authorization string
}

type Accounts struct {
	WithBalance   bool
	Date          string
	Authorization string
	RequestID     string
	ProcessID     string
}

type AccountList struct {
	List []Account `json:"account_list"`
}

type Account struct {
	ID             string               `json:"id"`
	Currency       string               `json:"currency"`
	Product        string               `json:"product"`
	AccountType    string               `json:"account_type"`
	IBAN           string               `json:"iban"`
	BIC            string               `json:"bic"`
	BBAN           string               `json:"bban"`
	ClearingNumber string               `json:"clearingnumber"`
	AccountNumber  string               `json:"account_number"`
	Balances       []map[string]Balance `json:"balances"`
}

type Balance struct {
	Amount Amount `json:"amount"`
	Date   string `json:"date"`
}

type Amount struct {
	Currency string  `json:"currency"`
	Content  float64 `json:"content"`
}

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

// GetAccounts ...
// This lists all payment accounts on the user
func GetAccounts(acc *Accounts) (*AccountList, error) {
	payload, err := http.NewRequest("GET", accounts, nil)
	payload.Header.Add("Authorization", acc.Authorization)
	payload.Header.Add("Process-ID", acc.ProcessID)
	payload.Header.Add("Request-ID", acc.RequestID)
	payload.Header.Add("Date", acc.Date)

	req, err := client.Do(payload)

	if err != nil {
		fmt.Println("Error occured while trying to fetch from", accounts)
	}

	defer req.Body.Close()
	decoder := json.NewDecoder(req.Body)

	if err != nil {
		return nil, err
	}

	accountList := AccountList{}
	err = decoder.Decode(&accountList)

	if err != nil {
		return nil, err
	}

	return &accountList, nil
}

// GetAccount ...
// This lists a single payment account on the user
func GetAccount(acc *SingleAccount) (*Account, error) {
	payload, err := http.NewRequest("GET", account+acc.ID+"?bic="+acc.BIC, nil)
	payload.Header.Add("Process-ID", acc.ProcessID)
	payload.Header.Add("Request-ID", acc.RequestID)
	payload.Header.Add("Authorization", acc.Authorization)

	req, err := client.Do(payload)

	if err != nil {
		fmt.Println("Error occured while trying to fetch from", account)
	}

	defer req.Body.Close()
	decoder := json.NewDecoder(req.Body)

	if err != nil {
		return nil, err
	}

	account := Account{}
	err = decoder.Decode(&account)

	if err != nil {
		return nil, err
	}

	return &account, nil
}

// GetTransactions ...
// Fetch transactions from the account
func GetTransactions(acc *SingleAccount) (*TransactionList, error) {
	return nil, nil
}
