package main

import (
	"fmt"
)

// Car is a car
type Car struct {
	VehicleInfo
}

// Type returns the car type
func (c Car) Type() VehicleType {
	return CarType
}

// Park will find a parking spot for a car
func (c Car) Park(pl *Lot) (ok bool, ps ParkingSpot) {
	const neededSpotsForCar = 1

	if ok, ps = pl.FindCompactSpot(neededSpotsForCar); ok {
	} else if ok, ps = pl.FindLargeSpot(neededSpotsForCar); ok {
	} else {
		fmt.Println("No Spot Availabel for a Car")
	}
	return
}

//NewCar returns an instance of a new Car
func NewCar() *Car {
	return &Car{VehicleInfo{NewPlate(), CarType}}
}
