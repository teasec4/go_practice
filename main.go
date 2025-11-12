package main

import (
	"fmt"
	"first_l/server"
)


func main() {
	// b := NewBalance(100)
	
	// PrintAccountInfo(b)

	// var dic map[int]string = map[int]string{1: "one", 2: "two"}
	// fmt.Println(dic[1])

	server.Start()
}


type Account interface {
    Withdraw(sum int) error
    Balance() int
}

type Balance struct {
	amount int
}

func (b *Balance) Balance() int {
	return b.amount
}

func (b *Balance) Withdraw(sum int) error{
	if sum <= 0 {
		return fmt.Errorf("sum must be positive")
	}
	if b.amount < sum{
		return  fmt.Errorf("insufficient funds")
	}
	b.amount -= sum
	return nil
}

func NewBalance(amount int) *Balance {
	return &Balance{amount: amount}
}

func PrintAccountInfo(a Account){
	fmt.Println("Balance:", a.Balance())
}