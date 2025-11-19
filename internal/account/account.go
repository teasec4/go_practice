// Package account provides account management functionality
package account

import "errors"

// Account defines the interface for bank account operations
type Account interface {
	Deposit(amount int) error
	Withdraw(amount int) error
	Balance() int
}

// impl is the concrete implementation of Account interface
type impl struct {
	balance int
}

// New creates a new account with the given initial balance
func New(initial int) Account {
	return &impl{balance: initial}
}

// Deposit adds money to the account
func (a *impl) Deposit(amount int) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	a.balance += amount
	return nil
}

// Balance returns the current account balance
func (a *impl) Balance() int {
	return a.balance
}

// Withdraw removes money from the account
func (a *impl) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("insufficient balance")
	}
	a.balance -= amount
	return nil
}
