// Create an interface called Car with methods Start(), Accelerate(), and Stop(). Create two structs, GasolineCar and
// ElectricCar, that implement the Car interface. GasolineCar should have a FuelLevel field and Accelerate() should
// decrease the fuel level. ElectricCar should have a BatteryLevel field and Accelerate() should decrease the battery level.
package main

import "fmt"

type car interface { // interface will applicable only all methods are used in program
	Start()
	Accelerate()
	Stop()
}

type GaslineCar struct {
	FuelLevel float64
	initiate  string
	stopp     string
}

type ElectricCar struct {
	BatteryLevel float64
	fueled       string
	drained      string
}

func (g *GaslineCar) Accelerate() { //(acc< fl)
	g.FuelLevel = 0.5

}
func (g *GaslineCar) Start() {
	g.initiate = "start"

}

func (g *GaslineCar) Stop() {
	g.stopp = "stop"
}

func (e *ElectricCar) Accelerate() {
	e.BatteryLevel = 2.5
}
func (e *ElectricCar) Start() {
	e.fueled = "started"
}
func (e *ElectricCar) Stop() {
	e.drained = "stopped"
}

//define function that takes a car nterfaces as a parameter and calls methods in it

func model(t car) {
	t.Accelerate()
	t.Start()
	t.Stop()

}

func main() {
	carr := &GaslineCar{FuelLevel: 10.5}
	c := GaslineCar{initiate: "started"}
	d := GaslineCar{stopp: "stop"}

	// var f car
	// f = &carr
	// fmt.Println(f)

	carr.Accelerate()
	c.Start()
	d.Stop()

	fmt.Println(carr)
	fmt.Println(c)
	fmt.Println(d)
	Electcar := ElectricCar{BatteryLevel: 10.5}
	e := ElectricCar{fueled: "started"}
	f := ElectricCar{drained: "stopped"}
	Electcar.Accelerate()
	e.Start()
	f.Stop()

	fmt.Println(Electcar)
	fmt.Println(e)
	fmt.Println(f)

	// //call the model function with GaslineCar and ElectricCar structs
	model(carr)
	// model (c)
	// model (d)
	// model(Electcar)

}
