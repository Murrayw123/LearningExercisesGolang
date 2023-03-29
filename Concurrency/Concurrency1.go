package Concurrency

import (
	"fmt"
	"sync"
)

type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64) (float64, error)
	Balance() float64
}

type BankAccount struct {
	balance float64
	lock    sync.Mutex
	Account
}

func (b *BankAccount) Deposit(amount float64) {
	b.lock.Lock()
	b.balance += amount
	b.lock.Unlock()
}

func (b *BankAccount) Withdraw(amount float64) (float64, error) {
	b.lock.Lock()
	b.balance -= amount
	b.lock.Unlock()
	return b.balance, nil
}

func (b *BankAccount) Balance() float64 {
	return b.balance
}

func ExerciseOne() {
	bankAccount := BankAccount{}
	wg := sync.WaitGroup{}

	deposit := func(amount float64) {
		bankAccount.Deposit(amount)
		wg.Done()
	}

	withdraw := func(amount float64) {
		_, err := bankAccount.Withdraw(amount)
		if err != nil {
			fmt.Println(err)
		}
		wg.Done()
	}

	for i := 0; i < 5; i++ {
		wg.Add(4)
		go deposit(100)
		go deposit(50)
		go deposit(20)
		go withdraw(100)
	}

	wg.Wait()
	fmt.Println("Balance: ", bankAccount.Balance())
}
