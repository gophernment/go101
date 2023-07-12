package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	account := NewAccount("go@gopher.com", time.Now().Unix())
	fmt.Printf("Email: %s\nCreated At: %s\n", account.Email, account.CreatedDate)
}

type Account struct {
	Email       string
	CreatedDate time.Time
}

func NewAccount(email string, CreatedDate int64) Account {
	return Account{}
}

func jsonString(account Account) (string, error) {
	b, err := json.Marshal(&account)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

var jsonExample = `{
    "email": "yod@go.dev",
    "created_date": "14-01-2022 09:10:11"
}
`
