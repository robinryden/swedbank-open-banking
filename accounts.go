package swedbank

import (
	"fmt"
	"net/http"
)

var (
	client   = &http.Client{}
	accounts = "https://psd2.api.swedbank.com/sandbox/v1/accounts"
)

type Accounts struct {
	BIC           string
	WithBalance   bool
	Date          string
	Authorization string
	RequestID     string
	ProcessID     string
}

// This lists all payment accounts on the user
func Get(acc *Accounts) int {
	payload, err := http.Get(accounts)
	payload.Header.Add("BIC", acc.BIC)
	payload.Header.Add("Authorization", acc.Authorization)
	payload.Header.Add("Process-ID", acc.ProcessID)
	payload.Header.Add("Request-ID", acc.RequestID)
	payload.Header.Add("Date", acc.Date)

	if err != nil {
		fmt.Println("Error occured while trying to fetch from", accounts)
	}

	defer payload.Body.Close()

	return payload.StatusCode
}
