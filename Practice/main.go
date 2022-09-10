package main

import "fmt"

func main() {
	b1 := BMW{distance: 100, fuel: 10, averagspeed: "58"}
	a1 := Audi{distance: 100, fuel: 10}
	person := []Vehicle{b1, a1}
	totalMilege(person)
}

type Vehicle interface {
	Mileage() float64
}

type BMW struct {
	distance    float64
	fuel        float64
	averagspeed string
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
func totalMilege(vehicle []Vehicle) {
	tm := 0.0
	for _, v := range vehicle {
		tm += v.Mileage()
	}
	fmt.Println("Total Mileage: ", tm)
}
