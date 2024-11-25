package main

import (
	"DebitAuthorizer/models"
	"fmt"
	"time"
)

func main() {
	var ts = &models.TransactionSettings{AmountLimit: 1000, PeriodLimit: 5 * time.Minute}
	th := models.NewTransactionHistory(ts)

	now := time.Now()
	fmt.Printf("Now: %s\n", now.Format(time.RFC3339))

	ExecuteCustomer1(th, now)

	fmt.Println("--------------------")

	ExecuteCustomer2(th, now)

}

func ExecuteCustomer1(th *models.TransactionHistory, now time.Time) {
	t1 := models.NewTransaction(20, now.Add(-5*time.Minute), 200)
	t2 := models.NewTransaction(20, now.Add(-4*time.Minute), 200)
	t3 := models.NewTransaction(20, now.Add(-3*time.Minute), 200)
	t4 := models.NewTransaction(20, now.Add(-2*time.Minute), 199.98)
	t5 := models.NewTransaction(20, now.Add(-1*time.Minute), 200.02)
	t6 := models.NewTransaction(20, now.Add(1*time.Minute), 200)
	//t7 := models.NewTransaction(20, now.Add(1*time.Minute), 0)

	showTransaction(th.AddTransaction(t1))
	showTransaction(th.AddTransaction(t2))
	showTransaction(th.AddTransaction(t3))
	showTransaction(th.AddTransaction(t4))
	showTransaction(th.AddTransaction(t5))
	showTransaction(th.AddTransaction(t6))
	//showTransaction(th.AddTransaction(t7))
}

func ExecuteCustomer2(th *models.TransactionHistory, now time.Time) {
	t1 := models.NewTransaction(10, now.Add(-5*time.Minute), 500)
	t2 := models.NewTransaction(10, now.Add(-4*time.Minute), 499.99)
	t3 := models.NewTransaction(10, now.Add(-3*time.Minute), 0.02)
	t4 := models.NewTransaction(10, now.Add(-3*time.Minute), 0.01)
	t5 := models.NewTransaction(10, now.Add(2*time.Minute), 0.01)
	t6 := models.NewTransaction(10, now.Add(3*time.Minute), 0.01)

	showTransaction(th.AddTransaction(t1))
	showTransaction(th.AddTransaction(t2))
	showTransaction(th.AddTransaction(t3))
	showTransaction(th.AddTransaction(t4))
	showTransaction(th.AddTransaction(t5))
	showTransaction(th.AddTransaction(t6))
}

func showTransaction(th *models.Transaction, err error) {
	if err == nil {
		fmt.Printf("\033[32mTransaction success %s\033[0m\n", th)
		return
	}

	fmt.Printf("\033[31mTransaction failed %s --> Cause: %s\033[0m\n", th, err)
}

// API REST:

//package main
//
//import (
//	"DebitAuthorizer/models"
//	"encoding/json"
//	"fmt"
//	"net/http"
//	"time"
//)
//
//var ts = &models.TransactionSettings{AmountLimit: 1000, PeriodLimit: 5 * time.Minute}
//var th = models.NewTransactionHistory(ts)
//
//func main() {
//	http.HandleFunc("/transaction", transactionHandler)
//	fmt.Println("Server started at :8080")
//	_ = http.ListenAndServe(":8080", nil)
//}
//
//func transactionHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodPost {
//		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
//		return
//	}
//
//	var req models.Transaction
//	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
//		http.Error(w, "Invalid request body", http.StatusBadRequest)
//		return
//	}
//
//	transaction := models.NewTransaction(req.CustomerID, req.DateTime, req.Amount)
//	_, err := th.AddTransaction(transaction)
//
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	w.WriteHeader(http.StatusCreated)
//}

//POST /transaction HTTP/1.1
//Content-Type: application/json
//User-Agent: insomnia/10.1.1
//Host: localhost:8080
//Content-Length: 74
//
//{
//	"customer_id": 1,
//	"amount": 200,
//	"datetime": "2025-01-01T00:00:00Z"
//}
