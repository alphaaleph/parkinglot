package main

import (
	"fmt"
)

func main() {

	//build the lot
	MyLot := new(Lot)
	BuildLot(MyLot)

	//check available spaces
	MyLot.AvailableSpots()

	// 3 motorcycles came to park
	for n := 0; n < 3; n++ {
		MyLot.FillSpot(NewMotorcycle())
	}
	MyLot.AvailableSpots()

	// 1 bus came to park
	MyLot.FillSpot(NewBus())
	MyLot.AvailableSpots()

	// 26 cars came to park
	for n := 0; n < 26; n++ {
		MyLot.FillSpot(NewCar())
	}
	MyLot.AvailableSpots()

	// 5 motorcycles came to park
	for n := 0; n < 5; n++ {
		MyLot.FillSpot(NewMotorcycle())
	}
	MyLot.AvailableSpots()

	// 1 bus came to park
	MyLot.FillSpot(NewBus())
	MyLot.AvailableSpots()

	//12 vehicles left
	for n := 0; n < 12; n++ {
		MyLot.Leave()
	}
	fmt.Println("12 vehicles left")
	MyLot.AvailableSpots()
}
