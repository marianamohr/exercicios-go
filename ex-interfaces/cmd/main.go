package main

import (
	"fmt"
	"reflect"
)

type Veiculo interface {
	TurnOn() bool
	TurnOff() bool
	Move(direction int, speed int) bool
	Stop() bool
}

type Gas interface {
	FillUp() bool
}

type Car struct {
	model string
}

type Truck struct {
	model string
}

func (c Car) TurnOn() bool {
	fmt.Println(" o Carro " + c.model + " está ligando")
	return true
}

func (c Car) TurnOff() bool {
	fmt.Println(" o Carro " + c.model + " está desligando")
	return true
}

func (c Car) Move(direction int, speed int) bool {
	fmt.Println(" o Carro " + c.model + " está andando")
	return true
}

func (c Car) Stop() bool {
	fmt.Println(" o Carro " + c.model + " está Parando")
	return true
}
func (c Car) FillUp() bool {
	fmt.Println("Enchendo o tanque do carro " + c.model)
	return true
}

func (t Truck) TurnOn() bool {
	fmt.Println(" o Carro " + t.model + " está ligando")
	return true
}

func (t Truck) TurnOff() bool {
	fmt.Println(" o Carro " + t.model + " está desligando")
	return true
}

func (t Truck) Move(direction int, speed int) bool {
	fmt.Println(" o Carro " + t.model + " está andando")
	return true
}

func (t Truck) Stop() bool {
	fmt.Println(" o Carro " + t.model + " está Parando")
	return true
}

func main() {
	fmt.Println("Inicio")
	defer fmt.Println("fim")
	newCar := Car{"Mustang"}
	fmt.Println(newCar, reflect.TypeOf(newCar))
	newCar.TurnOn()
	newCar.Move(1, 1)
	newCar.Stop()
	newCar.TurnOff()

	var newVeiculo Veiculo
	fmt.Println(newVeiculo, reflect.TypeOf(newVeiculo))

	newVeiculo = Car{"Dodge"}
	fmt.Println(newVeiculo, reflect.TypeOf(newVeiculo))
	newVeiculo.Move(1, 1)
	obj, ok := reflect.TypeOf(newVeiculo).MethodByName("FillUp")
	fmt.Println(ok, "testr")
	fmt.Println(obj, "obj")

	value, ok := newVeiculo.(Gas)
	fmt.Println(value, ok)

	var xablau interface{}
	xablau = Car{"Fusca"}
	fmt.Println(xablau)

}
