package models

import (
	"fmt"
	"time"
)

type Transaction struct {
	CustomerID int       `json:"customer_id"`
	DateTime   time.Time `json:"datetime"`
	Amount     float64   `json:"amount"`
	Succeeded  bool      `json:"succeeded"`
}

func NewTransaction(customerID int, dateTime time.Time, amount float64) *Transaction {
	return &Transaction{CustomerID: customerID, DateTime: dateTime, Amount: amount}
}

func (t *Transaction) String() string {
	return fmt.Sprintf("Transaction{CustomerID: %d, DateTime: %s, Amount: %.2f, Succeeded: %t}",
		t.CustomerID,
		t.DateTime.Format(time.RFC3339),
		t.Amount,
		t.Succeeded)
}
