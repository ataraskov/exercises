// Package account provides a solution to to Bank Account on Exercism's Go Track
package account

import "sync"

// Account represents a bank account
type Account struct {
	mu      sync.RWMutex
	closed  bool
	balance int64
}

// Open returns an Account with the specified balance
func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{balance: amount}
}

// Balance retunrs account's balance or false if account is closed
func (a *Account) Balance() (int64, bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if a.closed {
		return 0, false
	}
	return a.balance, true
}

// Deposit adjusts balance by amouny provided and false if operation is not possible
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.closed || a.balance+amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}

// Close markes account as closed and false is closed already
func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.closed {
		return 0, false
	}
	a.closed = true
	payout := a.balance
	a.balance = 0
	return payout, true
}
