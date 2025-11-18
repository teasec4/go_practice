package main
import (
	"errors"
)

// interface
type Account interface{
	Deposit(amount int) error
	Withdraw(amount int) error
	Balance() int
}

// Struct
type AccountImpl struct{
	balance int
}

// create new Account with initial amount
func NewAccount(initial int) Account{
	return  &AccountImpl{balance: initial}
}

// deposit
func (a *AccountImpl) Deposit(amount int) error{
	if amount <= 0 {
		return errors.New("Deposit amount must be positive")
	}
	a.balance += amount
	return nil
}

// check balance
func (a *AccountImpl) Balance() int{
	return a.balance
}

// withdraw
func (a *AccountImpl) Withdraw(amount int) error{
	if a.balance < amount{
		return errors.New("Insufficient balance")
	}
	a.balance -= amount
	return nil
}