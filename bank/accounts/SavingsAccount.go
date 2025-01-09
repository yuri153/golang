package accounts

import "bank/holders"

type SavingsAccount struct {
	Holder                                 holders.Holder
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (account *SavingsAccount) Withdraw(value float64) (string, float64, bool) {
	canWithdraw := value > 0 && value <= account.balance

	if canWithdraw {
		account.balance -= value

		return "Saque realizado com sucesso!", account.balance, true
	} else {
		return "Saldo insuficiente.", account.balance, false
	}
}

func (account *SavingsAccount) Deposit(value float64) (string, float64) {
	canDeposit := value > 0

	if canDeposit {
		account.balance += value

		return "Depósito realizado com sucesso!", account.balance
	} else {
		return "Depósito não pode ser menor que 0.", account.balance
	}
}
