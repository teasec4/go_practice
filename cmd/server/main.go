package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/teasec4/go_practice/internal/bank"
	"github.com/teasec4/go_practice/internal/handler"
)

func main() {
	port := flag.String("port", "8080", "Server port")
	flag.Parse()

	// Initialize bank
	b := bank.New()

	// Setup routes
	http.HandleFunc("/deposit", handler.Deposit(b))
	http.HandleFunc("/withdraw", handler.Withdraw(b))
	http.HandleFunc("/balance", handler.Balance(b))

	// Start server
	addr := ":" + *port
	fmt.Printf("Server starting on %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
