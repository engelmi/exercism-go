package account

import "sync"

type Account struct {
	closed bool

	currentBalance int64
	syncMutex      *sync.Mutex
}

func Open(initialBalance int64) *Account {
	if initialBalance < 0 {
		return nil
	}

	acc := &Account{
		closed:         false,
		currentBalance: initialBalance,
		syncMutex:      &sync.Mutex{},
	}

	return acc
}

func (acc *Account) Deposit(amount int64) (int64, bool) {
	acc.syncMutex.Lock()
	if acc.closed {
		acc.syncMutex.Unlock()
		return 0, false
	}
	balance := acc.currentBalance + amount
	if balance < 0 {
		acc.syncMutex.Unlock()
		return 0, false
	}
	acc.currentBalance = balance
	acc.syncMutex.Unlock()
	return balance, true
}

func (acc *Account) Balance() (int64, bool) {
	acc.syncMutex.Lock()
	if acc.closed {
		acc.syncMutex.Unlock()
		return 0, false
	}
	balance := acc.currentBalance
	acc.syncMutex.Unlock()
	return balance, true
}

func (acc *Account) Close() (int64, bool) {
	acc.syncMutex.Lock()
	if acc.closed {
		acc.syncMutex.Unlock()
		return 0, false
	}
	acc.closed = true
	balance := acc.currentBalance
	acc.syncMutex.Unlock()
	return balance, true
}
