package main

import (
	"fmt"
)

// Motorcycle is a motorcycle
type Motorcycle struct {
	VehicleInfo
}

// Type returns the motorcycle type
func (m Motorcycle) Type() VehicleType {
	return MotorcycleType
}

// Park will find a parking spot for a motorcycle
func (m Motorcycle) Park(pl *Lot) (ok bool, ps ParkingSpot) {
	const neededSpotsForMotorcycle = 1

	if ok, ps = pl.FindMotorcycleSpot(neededSpotsForMotorcycle); ok {
	} else if ok, ps = pl.FindCompactSpot(neededSpotsForMotorcycle); ok {
	} else if ok, ps = pl.FindLargeSpot(neededSpotsForMotorcycle); ok {
	} else {
		fmt.Println("No Spot Available for a Motorcycle")
	}
	return
}

//NewMotorcycle returns an instance of a new Motorcyfle
func NewMotorcycle() *Motorcycle {
	return &Motorcycle{VehicleInfo{NewPlate(), MotorcycleType}}
}
