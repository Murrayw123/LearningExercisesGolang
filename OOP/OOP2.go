package OOP

import (
	"fmt"
)

type Vehicle interface {
	Start() string
	Stop() string
	GetStatus() string
}

type VehicleImpl struct {
	status string
}

func newVehicle() VehicleImpl {
	return VehicleImpl{
		status: "Stopped",
	}
}

func (v *VehicleImpl) GetStatus() string {
	return v.status
}

func (v *VehicleImpl) Start() string {
	v.status = "Started"
	return "Starting Vehicle"
}

func (v *VehicleImpl) Stop() string {
	v.status = "Stopped"
	return "Stopping Vehicle"
}

type Car struct {
	VehicleImpl
}

type Motorcycle struct {
	VehicleImpl
}

type Truck struct {
	VehicleImpl
}

func OOPExerciseTwo() {

	car := Car{newVehicle()}
	motorcycle := Motorcycle{newVehicle()}

	fmt.Println("car status", car.GetStatus())
	fmt.Println("car starting", car.Start())
	fmt.Println("car status", car.GetStatus())

	fmt.Println("motorcycle status", motorcycle.GetStatus())
	fmt.Println("motorcycle starting", motorcycle.Start())
	fmt.Println("motorcycle status", motorcycle.GetStatus())

}
