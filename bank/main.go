package main

import (
	"fmt"
	"golang/bank/accounts"
	"golang/bank/holders"
)

func main() {
	account := accounts.CheckingAccount{Holder: holders.Holder{Name: "Yuri"}, AgencyNumber: 589, AccountNumber: 123456}

	fmt.Println(account.Withdraw(50))
}
