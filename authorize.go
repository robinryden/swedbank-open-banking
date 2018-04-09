package swedbank

import (
	"fmt"
	"net/http"
)

const (
	authorizeAPI = "https://psd2.api.swedbank.com/psd2/authorize"
)

type Authorize struct {
	BIC              string
	TPPTransactionID string
	TPPRequestID     string
	ClientID         string
	ClientSecret     string
	RedirectURI      string
	ResponseType     string
	Scope            string
}

func Auth(auth *Authorize) (string, error) {
	payload, err := http.NewRequest("POST", authorizeAPI+"?bic="+auth.BIC+"&client_id="+auth.ClientID+"&client_secret="+auth.ClientSecret+"&redirect_uri="+auth.RedirectURI+"&response_type="+auth.ResponseType+"&scope="+auth.Scope, nil)
	payload.Header.Add("TPP-Transaction-ID", auth.TPPTransactionID)
	payload.Header.Add("TPP-Request-ID", auth.TPPRequestID)

	req, err := client.Do(payload)

	if err != nil {
		fmt.Println("Error occured while trying to fetch from", err)
	}

	defer req.Body.Close()
	// decoder := json.NewDecoder(req.Body)

	if err != nil {
		return "", err
	}

	fmt.Println(req.StatusCode)
	fmt.Println(req.Request.URL)

	// body, err := ioutil.ReadAll(req.Body)
	// if err != nil {
	// 	//
	// }
	// fmt.Println(string(body))

	// transaction := TransactionList{}
	// err = decoder.Decode(&transaction)

	if err != nil {
		return "", err
	}

	return "", nil
}
