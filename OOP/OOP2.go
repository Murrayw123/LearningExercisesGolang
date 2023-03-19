package OOP

import "go/types"

type Vehicle interface {
	Start() types.Nil
	Stop() types.Nil
}

type Car struct{}

type Motorcycle struct{}

type Truck struct{}

func (c Car) Start() types.Nil {
	return c.Start()
}

func (c Car) Stop() types.Nil {
	return c.Stop()
}

func (m Motorcycle) Start() types.Nil {
	return m.Start()
}

func (m Motorcycle) Stop() types.Nil {
	return m.Stop()
}

func (t Truck) Start() types.Nil {
	return t.Start()
}

func (t Truck) Stop() types.Nil {
	return t.Stop()
}

func OOPExerciseTwo() {
	car := Car{}
	motorcycle := Motorcycle{}
	truck := Truck{}

	car.Start()
	car.Stop()

	motorcycle.Start()
	motorcycle.Stop()

	truck.Start()
}
