package OOP

import "testing"

func TestOOPExerciseThree_get_and_set_name(t *testing.T) {
	person := Person{}
	person.SetName("John")
	name := person.GetName()
	if name != "John" {
		t.Errorf("Expected name to be John, got %s", name)
	}
}

func TestOOPExerciseThree_get_and_set_age(t *testing.T) {
	person := Person{}
	person.SetAge(25)
	age := person.GetAge()
	if age != 25 {
		t.Errorf("Expected age to be 25, got %d", age)
	}
}
