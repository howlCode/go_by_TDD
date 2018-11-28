package main

import (
	"errors"
	"fmt"
)

type BitCoin int

type Stringer interface {
	String() string
}

func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance BitCoin
}

func (w *Wallet) Balance() BitCoin {
	return w.balance
}

func (w *Wallet) Deposit(amount BitCoin) {
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount BitCoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func main() {
	fmt.Println("Pointers and Errors")
}
