package OOP

type Person struct {
	name string
	age  int
}

func (p *Person) validateName(name string) bool {
	return len(name) > 0
}

func (p *Person) validateAge(age int) bool {
	return age > 0
}

func (p *Person) SetName(name string) {
	if p.validateName(name) {
		p.name = name
	}
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) SetAge(age int) {
	if p.validateAge(age) {
		p.age = age
	}
}

func (p *Person) GetAge() int {
	return p.age
}

func OOPExerciseThree() {
	person := Person{}
	person.SetName("John")
	person.SetAge(25)
	name := person.GetName()
	age := person.GetAge()
	if name != "John" {
		panic("Expected name to be John")
	}
	if age != 25 {
		panic("Expected age to be 25")
	}
}
