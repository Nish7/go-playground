package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(val Bitcoin) {
	w.balance += val
}

func (w *Wallet) Withdraw(val Bitcoin) error {
	if val > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= val
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
