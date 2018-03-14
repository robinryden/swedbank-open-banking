package swedbank

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	client   = &http.Client{}
	accounts = "https://psd2.api.swedbank.com/sandbox/v1/accounts/?BIC=SANDSESS"
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
func Get(acc *Accounts) {
	serilizedPayload, err := json.Marshal(acc)
	if err != nil {
		fmt.Println(err)
	}

	payload, err := http.NewRequest("GET", accounts, bytes.NewBuffer(serilizedPayload))
	resp, err := client.Do(payload)

	if err != nil {
		fmt.Println("Error occured while trying to fetch from", accounts)
	}

	defer resp.Body.Close()
}
