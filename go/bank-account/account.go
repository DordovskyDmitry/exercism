package account

import "sync"

const testVersion = 1

type Account struct {
	balance int64
	closed  bool
	mux     sync.Mutex
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{balance: initialDeposit, closed: false}
}

func (a *Account) Close() (payout int64, ok bool) {
	a.mux.Lock()
	defer a.mux.Unlock()

	if a.closed {
		return 0, false
	}
	a.closed = true
	return a.balance, true
}

func (a *Account) Balance() (balance int64, ok bool) {
	if a.closed {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mux.Lock()
	defer a.mux.Unlock()

	if a.closed || a.balance+amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}
