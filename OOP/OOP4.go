package OOP

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	balance      float64
	balanceState BalanceStates
}

type BalanceStates int

const (
	healthy BalanceStates = iota
	overdrawn
)

func (b *BankAccount) Deposit(amount float64) {
	b.balance += amount
}

func (b *BankAccount) Withdraw(amount float64) (float64, error) {
	balanceState, err := b.validateBalance(amount)
	if err != nil {
		return 0, err
	}
	b.balance -= amount
	b.balanceState = balanceState
	return b.balance, nil
}

func (b *BankAccount) Balance() float64 {
	return b.balance
}

func (b *BankAccount) validateBalance(withdrawAmount float64) (BalanceStates, error) {
	if (b.balance-withdrawAmount < 0) && (b.balance-withdrawAmount > -100) {
		return overdrawn, nil
	} else if (b.balance-withdrawAmount < 0) && (b.balance-withdrawAmount < -100) {
		return overdrawn, errors.New("Overdrawn, can not withdraw further")
	} else {
		return healthy, nil
	}
}

func OOPExerciseFour() {
	bankAccount := BankAccount{}
	bankAccount.Deposit(100)
	balance, err := bankAccount.Withdraw(50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Balance: ", balance)

}
