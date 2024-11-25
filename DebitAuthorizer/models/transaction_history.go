package models

import (
	"fmt"
	"time"
)

type TransactionSettings struct {
	AmountLimit float64
	PeriodLimit time.Duration
}

type TransactionHistory struct {
	Transactions        []Transaction
	TransactionSettings *TransactionSettings
}

func NewTransactionHistory(transactionSettings *TransactionSettings) *TransactionHistory {
	return &TransactionHistory{
		Transactions:        []Transaction{},
		TransactionSettings: transactionSettings,
	}
}

func (th *TransactionHistory) AddTransaction(currentTransaction *Transaction) (*Transaction, error) {
	isLimitReached, err := th.isTransactionLimitReached(currentTransaction)

	if isLimitReached {
		currentTransaction.Succeeded = false
		return currentTransaction, err
	}

	currentTransaction.Succeeded = true
	th.Transactions = append(th.Transactions, *currentTransaction)
	return currentTransaction, nil
}

func (th *TransactionHistory) isTransactionLimitReached(currentTransaction *Transaction) (bool, error) {
	latestTransaction := th.getLatestTransactionByCustomerID(currentTransaction.CustomerID)

	if latestTransaction == nil {
		return false, nil
	}

	totalAmount := th.totalAmountByCustomerID(currentTransaction.CustomerID) + currentTransaction.Amount

	isTotalAmountGreaterThanLimit := totalAmount > th.TransactionSettings.AmountLimit

	isWithinPeriod := currentTransaction.DateTime.Sub(latestTransaction.DateTime) <= th.TransactionSettings.PeriodLimit

	if isTotalAmountGreaterThanLimit && isWithinPeriod {
		err := fmt.Errorf("transaction limit reached: allowed amount for new transactions is %.2f or less within a period of %s",
			th.TransactionSettings.AmountLimit,
			th.TransactionSettings.PeriodLimit)

		return true, err
	}

	return false, nil
}

func (th *TransactionHistory) getLatestTransactionByCustomerID(customerID int) *Transaction {
	var latestTransaction *Transaction

	for _, t := range th.Transactions {
		if t.CustomerID == customerID {
			if latestTransaction == nil || t.DateTime.After(latestTransaction.DateTime) {
				latestTransaction = &t
			}
		}
	}

	return latestTransaction
}

func (th *TransactionHistory) totalAmountByCustomerID(customerID int) float64 {
	var total float64

	for _, t := range th.Transactions {
		if t.CustomerID == customerID {
			total += t.Amount
		}
	}

	return total
}
