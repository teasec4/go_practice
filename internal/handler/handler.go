// Package handler provides HTTP request handlers
package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/teasec4/go_practice/internal/bank"
)

type depositRequest struct {
	AccountID int `json:"account_id"`
	Amount    int `json:"amount"`
}

type withdrawRequest struct {
	AccountID int `json:"account_id"`
	Amount    int `json:"amount"`
}

type errorResponse struct {
	Error string `json:"error"`
}

type depositResponse struct {
	Message string `json:"message"`
	Balance int    `json:"balance"`
}

type balanceResponse struct {
	AccountID int `json:"account_id"`
	Balance   int `json:"balance"`
}

// Deposit handles POST /deposit requests
func Deposit(b *bank.Bank) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req depositRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorResponse{Error: "invalid request"})
			return
		}

		account, exists := b.GetAccount(req.AccountID)
		if !exists {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(errorResponse{Error: "account not found"})
			return
		}

		if err := account.Deposit(req.Amount); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(depositResponse{
			Message: "deposit successful",
			Balance: account.Balance(),
		})
	}
}

// Withdraw handles POST /withdraw requests
func Withdraw(b *bank.Bank) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req withdrawRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorResponse{Error: "invalid request"})
			return
		}

		account, exists := b.GetAccount(req.AccountID)
		if !exists {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(errorResponse{Error: "account not found"})
			return
		}

		if err := account.Withdraw(req.Amount); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(depositResponse{
			Message: "withdraw successful",
			Balance: account.Balance(),
		})
	}
}

// Balance handles GET /balance requests
func Balance(b *bank.Bank) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		accountIDStr := r.URL.Query().Get("account_id")
		if accountIDStr == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorResponse{Error: "account_id parameter required"})
			return
		}

		accountID, err := strconv.Atoi(accountIDStr)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorResponse{Error: "invalid account_id"})
			return
		}

		account, exists := b.GetAccount(accountID)
		if !exists {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(errorResponse{Error: "account not found"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(balanceResponse{
			AccountID: accountID,
			Balance:   account.Balance(),
		})
	}
}
