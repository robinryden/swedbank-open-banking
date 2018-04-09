# swedbank-open-banking
Go package for Swedbanks Open Banking API

### Documentation from Swedbanks homepage
[Swedbanks Open Banking page](https://www.swedbank.com/openbanking/)

[Developer Documentation for Swedbank Open Banking Sandbox (BETA)](https://www.swedbank.se/idc/groups/public/@i/@sc/@all/@kp/documents/regulation/cid_2441155.pdf)

### Good to know:
This is a work in progress package written in Go that uses the Swedbank Open Banking API.
All feedback is appreciated!

## Examples
### following examples are used inside Sandbox mode

#### Get all accounts from user

```go
accounts, err := swedbank.GetAccounts(&swedbank.Accounts{
  "SANDSESS",
  true,
  "Thu, 01 Dec 1994 16:00:00 GMT",
  "12345SGHDF",
  "AZXS3456",
  "Bearer dummyToken",
})

if err != nil {
  fmt.Println(err)
}

for _, account := range accounts.List {
  fmt.Println(account.ID)
}
```

#### Get specific account from user with account-id

```go
account, err := swedbank.GetAccount(&swedbank.SingleAccount{
  "AbcD1234eFgH568",
  "SANDSESS",
  true,
  "12345SGHDF",
  "AZXS3456",
  "Bearer dummyToken",
})

if err != nil {
  fmt.Println(err)
}

fmt.Println(account)
```

#### Get transactions list from specific account

```go
transactions, err := swedbank.GetTransactions(&swedbank.SingleAccount{
  "AbcD1234eFgH568",
  "SANDSESS",
  true,
  "12345SGHDF",
  "AZXS3456",
  "Bearer dummyToken",
})

if err != nil {
  fmt.Println(err)
}

for _, transaction := range transactions.List {
  fmt.Println(transaction)
}
```
