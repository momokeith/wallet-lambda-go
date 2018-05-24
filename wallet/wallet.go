package wallet

import (
	"encoding/json"
	"../rds"
)

type Wallet struct {
	Uuid string
	User_id string
}

type Wallets []Wallet

func List() Wallets {
	db := rds.GetDB()
	defer db.Close()
	var wallets []Wallet
	db.Find(&wallets)

	return wallets
}

func (wallets Wallets) Json() string {
	b, _ :=  json.Marshal(wallets)
	return string(b)
}