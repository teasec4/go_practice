package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Bank struct {
	accounts map[int]Account
	mu       sync.RWMutex
}

var bank = &Bank{
	accounts: make(map[int]Account),
}

func init() {
	// Create default account with ID 1
	bank.accounts[1] = NewAccount(1000)
}

func depositHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		AccountID int `json:"account_id"`
		Amount    int `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	bank.mu.Lock()
	account, exists := bank.accounts[req.AccountID]
	bank.mu.Unlock()

	if !exists {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	if err := account.Deposit(req.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Deposit successful",
		"balance": account.Balance(),
	})
}

func withdrawHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		AccountID int `json:"account_id"`
		Amount    int `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	bank.mu.Lock()
	account, exists := bank.accounts[req.AccountID]
	bank.mu.Unlock()

	if !exists {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	if err := account.Withdraw(req.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Withdraw successful",
		"balance": account.Balance(),
	})
}

func balanceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	accountIDStr := r.URL.Query().Get("account_id")
	if accountIDStr == "" {
		http.Error(w, "account_id query parameter required", http.StatusBadRequest)
		return
	}

	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		http.Error(w, "Invalid account_id", http.StatusBadRequest)
		return
	}

	bank.mu.RLock()
	account, exists := bank.accounts[accountID]
	bank.mu.RUnlock()

	if !exists {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"account_id": accountID,
		"balance":    account.Balance(),
	})
}

func StartServer(port string) {
	http.HandleFunc("/deposit", depositHandler)
	http.HandleFunc("/withdraw", withdrawHandler)
	http.HandleFunc("/balance", balanceHandler)

	fmt.Printf("Server starting on :%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
