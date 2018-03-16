package swedbank

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	client   = &http.Client{}
	accounts = "https://psd2.api.swedbank.com/sandbox/v1/accounts/?bic=SANDSESS"
)

type Accounts struct {
	WithBalance   bool
	Date          string
	Authorization string
	RequestID     string
	ProcessID     string
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

type AccountList struct {
	List []Account `json:"account_list"`
}

type Balance struct {
	Amount Amount `json:"amount"`
	Date   string `json:"date"`
}

type Amount struct {
	Currency string  `json:"currency"`
	Content  float64 `json:"content"`
}

// This lists all payment accounts on the user
func Get(acc *Accounts) (*AccountList, error) {
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

	test := AccountList{}
	err = decoder.Decode(&test)

	if err != nil {
		return nil, err
	}

	return &test, nil
}
