package main

import (
	"flag"
	"fmt"
)

func main() {
	mode := flag.String("mode", "server", "Choose mode: cli or server")
	port := flag.String("port", "8080", "Port for server mode")
	flag.Parse()

	if *mode == "server" {
		StartServer(*port)
	} else {
		runCLI()
	}
}

func runCLI() {
	fmt.Println("Hello! Welcome to Max Bank, How we can help you?")
	account := NewAccount(1000)
	fmt.Println("Balance:", account.Balance())

	for {
		fmt.Println("\nPlease enter a command:")
		fmt.Println("1. deposit")
		fmt.Println("2. withdraw")
		fmt.Println("3. balance")
		fmt.Println("4. exit")
		fmt.Print("Enter command: ")

		var cmd string
		fmt.Scan(&cmd)

		switch cmd {
		case "1":
			var amount int
			fmt.Print("Enter amount: ")
			fmt.Scan(&amount)
			if err := account.Deposit(amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful!")
			}
			fmt.Println("Balance:", account.Balance())
		case "2":
			var amount int
			fmt.Print("Enter amount: ")
			fmt.Scan(&amount)
			if err := account.Withdraw(amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdraw successful!")
			}
			fmt.Println("Balance:", account.Balance())
		case "3":
			fmt.Println("Balance:", account.Balance())
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}
