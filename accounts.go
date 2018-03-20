package swedbank

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	accountAPI = "https://psd2.api.swedbank.com/sandbox/v1/accounts/"
)

var (
	client = &http.Client{}
)

// Accounts
type SingleAccount struct {
	ID            string
	BIC           string
	WithBalance   bool
	RequestID     string
	ProcessID     string
	Authorization string
}

type Accounts struct {
	BIC           string
	WithBalance   bool
	Date          string
	RequestID     string
	ProcessID     string
	Authorization string
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

// GetAccounts ...
// This lists all payment accounts on the user
func GetAccounts(acc *Accounts) (*AccountList, error) {
	payload, err := http.NewRequest("GET", accountAPI+"/?bic="+acc.BIC, nil)
	payload.Header.Add("Authorization", acc.Authorization)
	payload.Header.Add("Process-ID", acc.ProcessID)
	payload.Header.Add("Request-ID", acc.RequestID)
	payload.Header.Add("Date", acc.Date)

	req, err := client.Do(payload)

	if err != nil {
		fmt.Println("Error occured while trying to fetch from", accountAPI)
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
	payload, err := http.NewRequest("GET", accountAPI+acc.ID+"?bic="+acc.BIC, nil)
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

	account := Account{}
	err = decoder.Decode(&account)

	if err != nil {
		return nil, err
	}

	return &account, nil
}
