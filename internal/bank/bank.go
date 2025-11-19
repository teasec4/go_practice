// Package bank provides bank management functionality
package bank

import (
	"sync"

	"github.com/teasec4/go_practice/internal/account"
)

// Bank manages multiple accounts
type Bank struct {
	accounts map[int]account.Account
	mu       sync.RWMutex
}

// New creates a new bank instance with a default account
func New() *Bank {
	b := &Bank{
		accounts: make(map[int]account.Account),
	}
	// Create default account with ID 1
	b.accounts[1] = account.New(1000)
	return b
}

// GetAccount retrieves an account by ID
func (b *Bank) GetAccount(id int) (account.Account, bool) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	acc, exists := b.accounts[id]
	return acc, exists
}

// LockAccount locks the account for write operations
func (b *Bank) LockAccount(id int) (account.Account, bool) {
	b.mu.Lock()
	defer b.mu.Unlock()
	acc, exists := b.accounts[id]
	return acc, exists
}
