package OOP

import "testing"

func TestOOPExerciseFour_get_and_set_happy_path_balance(t *testing.T) {
	bankAccount := BankAccount{}
	bankAccount.Deposit(100)
	balance := bankAccount.Balance()
	if balance != 100 {
		t.Errorf("Expected balance to be 100, got %f", balance)
	}
}

func TestOOPExerciseFour_should_be_able_to_overdraw(t *testing.T) {
	bankAccount := BankAccount{}
	bankAccount.Deposit(100)
	_, err := bankAccount.Withdraw(150)
	if err != nil {
		t.Errorf("Expected to be able to overdraw, got %s", err)
	}
}

func TestOOPExerciseFour_should_not_be_able_to_overdraw_more_than_100(t *testing.T) {
	bankAccount := BankAccount{}
	bankAccount.Deposit(100)
	_, err := bankAccount.Withdraw(250)
	if err == nil {
		t.Errorf("Expected to not be able to overdraw more than 100, got %s", err)
	}
}
