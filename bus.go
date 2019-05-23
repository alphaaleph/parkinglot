package main

import (
	"fmt"
)

//Bus is a bus
type Bus struct {
	VehicleInfo
}

// Type returns the bus type
func (b Bus) Type() VehicleType {
	return BusType
}

// Park will find a parking spot for a bus
func (b Bus) Park(pl *Lot) (ok bool, ps ParkingSpot) {

	const neededSpotsForBus = 5

	if ok, ps = pl.FindConsecutiveLargeSpot(neededSpotsForBus); ok {
	} else {
		fmt.Println("No Spot Available for a Bus")
	}
	return
}

//NewBus returns an instance of a new Bus
func NewBus() *Bus {
	return &Bus{VehicleInfo{NewPlate(), BusType}}
}
