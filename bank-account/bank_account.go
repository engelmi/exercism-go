package account

import "sync"

type Account struct {
	closed bool

	currentBalance int64
	sync.Mutex
}

func Open(initialBalance int64) *Account {
	if initialBalance < 0 {
		return nil
	}

	acc := &Account{
		closed:         false,
		currentBalance: initialBalance,
		Mutex:          sync.Mutex{},
	}

	return acc
}

func (acc *Account) Deposit(amount int64) (int64, bool) {
	acc.Lock()
	defer acc.Unlock()

	if acc.closed {
		return acc.currentBalance, !acc.closed
	}

	newBalance := acc.currentBalance + amount
	newBalanceOk := newBalance >= 0
	if newBalanceOk {
		acc.currentBalance = newBalance
	}

	return acc.currentBalance, !acc.closed && newBalanceOk
}

func (acc *Account) Balance() (int64, bool) {
	return acc.currentBalance, !acc.closed
}

func (acc *Account) Close() (int64, bool) {
	acc.Lock()
	defer acc.Unlock()

	if acc.closed {
		return 0, false
	}

	acc.closed = true
	payout := acc.currentBalance
	acc.currentBalance = 0

	return payout, true
}
