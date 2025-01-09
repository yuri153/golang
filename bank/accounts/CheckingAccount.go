package accounts

import "golang/bank/holders"

type CheckingAccount struct {
	Holder                      holders.Holder
	AgencyNumber, AccountNumber int
	balance                     float64
}

func (account *CheckingAccount) Withdraw(value float64) (string, float64, bool) {
	canWithdraw := value > 0 && value <= account.balance

	if canWithdraw {
		account.balance -= value

		return "Saque realizado com sucesso!", account.balance, true
	} else {
		return "Saldo insuficiente.", account.balance, false
	}
}

func (account *CheckingAccount) Deposit(value float64) (string, float64) {
	canDeposit := value > 0

	if canDeposit {
		account.balance += value

		return "Depósito realizado com sucesso!", account.balance
	} else {
		return "Depósito não pode ser menor que 0.", account.balance
	}
}

func (account *CheckingAccount) Transfer(value float64, accountToTransfer *CheckingAccount) bool {

	_, _, result := account.Withdraw(value)

	if result {
		accountToTransfer.Deposit(value)
	}

	return result
}

func (account *CheckingAccount) GetBalance() float64 {
	return account.balance
}
