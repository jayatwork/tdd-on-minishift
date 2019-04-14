package main

import (
	"errors"
	"fmt"
)

//Bitcoin represents a number of Bitcoins
type Bitcoin int

func (b Bitcoin) String() string {

	return fmt.Sprintf("%d BTC", b)
}

//Wallet stores the number of BTC someone owns
type Wallet struct {
	balance Bitcoin
}

//Deposit will add some Bitcoin to a wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//Withdram deducts some bitcoin from wallet, returning an error if cannot be performed
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return errors.New("uh oh")
	}

	w.balance -= amount
	return nil
}

//Balance returns the number of Bitcoin a wallet has
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
