package main

import (
	"bank/accounts"
	"bank/holders"
	"fmt"
)

func main() {
	account1 := accounts.CheckingAccount{Holder: holders.Holder{Name: "José"}, AgencyNumber: 589, AccountNumber: 123456}
	account1.Deposit(1000)

	account2 := accounts.CheckingAccount{Holder: holders.Holder{Name: "Maria"}, AgencyNumber: 589, AccountNumber: 123457}
	account2.Deposit(500)

	payBill(&account1, 100)
	fmt.Println(account1.GetBalance())

	account2.Withdraw(150)
	fmt.Println(account2.GetBalance())

	account1.Transfer(250, &account2)
	fmt.Println(account1.GetBalance())
	fmt.Println(account2.GetBalance())
}

type account interface {
	Withdraw(value float64) (string, float64, bool)
}

func payBill(account account, value float64) {
	account.Withdraw(value)
}
