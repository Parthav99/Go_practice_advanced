package main

import "fmt"

type MotorVehicle interface {
	Mileage() float64
}

type BMW struct {
	distance     float64
	fuel         float64
	averagespeed string
}

type Audi struct {
	distance float64
	fuel     float64
}

func (b BMW) Mileage() float64 {
	return b.distance / b.fuel
}

func (a Audi) Mileage() float64 {
	return a.distance / a.fuel
}

func totalMileage(m []MotorVehicle) {
	totalMileage := 0.0
	for _, v := range m {
		totalMileage = totalMileage + v.Mileage()
	}
	fmt.Println(totalMileage)
}

type emptyInterface interface{}

func checkType(i emptyInterface) {
	fmt.Printf("type %T, value %v\n", i, i)
}

func main() {

	b1 := BMW{
		distance:     239.44,
		fuel:         72.77,
		averagespeed: "55",
	}

	a1 := Audi{
		distance: 239.44,
		fuel:     72.77,
	}

	persons := []MotorVehicle{a1, b1}
	totalMileage(persons)

	str := "Hello World"
	checkType(str)

}
