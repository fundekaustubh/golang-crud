package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Account struct {
	FirstName             string
	LastName              string
	Email                 string
	AccountID             string
	Friends               []string
	PendingFriendRequests []string
}

func main() {
	accounts := []Account{}
	accountsHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			newAccount := Account{
				FirstName:             r.FormValue("firstName"),
				LastName:              r.FormValue("lastName"),
				Email:                 r.FormValue("email"),
				AccountID:             uuid.New().String(),
				Friends:               []string{},
				PendingFriendRequests: []string{},
			}
			accounts = append(accounts, newAccount)
			fmt.Fprintf(w, "Account created successfully!")
		} else if r.Method == http.MethodGet {
			fmt.Fprintf(w, "Accounts: %v", accounts)
		} else if r.Method == http.MethodDelete {
			accountID := r.FormValue("accountID")
			for i, account := range accounts {
				if account.AccountID == accountID {
					accounts = append(accounts[:i], accounts[i+1:]...)
					fmt.Fprintf(w, "Account deleted successfully!")
					return
				}
			}
			fmt.Fprintf(w, "Account not found!")
		} else {
			fmt.Fprintf(w, "Method not allowed!")
		}
	}
	http.HandleFunc("/accounts", accountsHandler)
	fmt.Println("Server is listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
