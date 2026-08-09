package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int
var ErrorInsuficientFunds error = errors.New("insufficient funds")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%dBTC", b)
}

type Wallet struct {
	balance Bitcoin 
}

func (w *Wallet) Deposit(quantity Bitcoin) {
	w.balance += quantity
}

func (w *Wallet) Withdraw(quantity Bitcoin) error {
	if quantity > w.balance{
		return ErrorInsuficientFunds 
	}
	w.balance -= quantity
	return nil
}

func (w *Wallet) Balance() Bitcoin{
	return w.balance
}

