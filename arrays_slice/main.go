// Create an interface called Car with methods Start(), Accelerate(), and Stop(). Create two structs, GasolineCar and
// ElectricCar, that implement the Car interface. GasolineCar should have a FuelLevel field and Accelerate() should
// decrease the fuel level. ElectricCar should have a BatteryLevel field and Accelerate() should decrease the battery level.
package main

import "fmt"

type car interface { // interface will applicable only all methods are used in program
	//Start()
	Acceleration() int
	//Stop()
}

type GaslineCar struct {
	velocity int
	time     int
}

type ElectricCar struct {
	velocity int
	time     int
}

func (g GaslineCar) Acceleration() int {
	return g.velocity / g.time
}

func (e ElectricCar) Acceleration() int {
	return e.velocity / e.time
}

// calculating total acceleration that both car  consumed

func totalacceleration(c []car) {
	ta := 0.0 //inital value
	for _, v := range c {
		ta = ta + float64(v.Acceleration())
	}

	fmt.Printf("total acceleration for both cars per month %f v/t", ta)

}

func main() {
	g1 := GaslineCar{velocity: 30, time: 30}
	e1 := ElectricCar{velocity: 60, time: 20}

	// initiaalising person who accessing acceleartaion of both cars

	person := []car{g1, e1}
	totalacceleration(person)

}

// func (g *GaslineCar) Accelerate() { //(acc< fl)
// 	g.FuelLevel = 0.5

// }
// func (g *GaslineCar) start() {
// 	g.initiate = "start"

// }

// func (g *GaslineCar) stop() {
// 	g.stopp = "stop"
// }

// func (e *ElectricCar) Accelerate() {
// 	e.BatteryLevel = 2.5
// }
// func (e *ElectricCar) start() {
// 	e.fueled = "started"
// }
// func (e *ElectricCar) stop() {
// 	e.drained = "stopped"
// }

// //define function that takes a car nterfaces as a parameter and calls methods in it

// func model(t car) {
// 	t.Accelerate()
// 	t.Start()
// 	t.Stop()

// }

// func main() {
// 	carr := GaslineCar{FuelLevel: 10.5}
// 	c := GaslineCar{initiate: "started"}
// 	d := GaslineCar{stopp: "stop"}

// 	// var f car
// 	// f = &carr
// 	// fmt.Println(f)

// 	carr.Accelerate()
// 	c.start()
// 	d.stop()

// 	fmt.Println(carr)
// 	fmt.Println(c)
// 	fmt.Println(d)
// 	Electcar := ElectricCar{BatteryLevel: 10.5}
// 	e := ElectricCar{fueled: "started"}
// 	f := ElectricCar{drained: "stopped"}

// 	Electcar.Accelerate()
// 	e.start()
// 	f.stop()

// 	fmt.Println(Electcar)
// 	fmt.Println(e)
// 	fmt.Println(f)

// 	// //call the model function with GaslineCar and ElectricCar structs
// 	// model(carr)
// 	// model(Electcar)

// }

//func main() {

// 	//a := [10]int{10, 30, 40, 50}
// 	a := [10]int{10, 20, 30, 90, 7, 8, 9}
// 	b := []string{"siri", "a", "b", "c"}
// 	// a[5] = 60
// 	// a[6] = 70
// 	print(a)

// 	//fmt.Println(a)

// func print(a [10]int) {

// 	for i, v := range a {
// 		// if 10 < 30 {
// 		// 	fmt.Println("10 is less than 30 in iteration")
// 		// }
// 		fmt.Println(i, v)
// 	}

// }
