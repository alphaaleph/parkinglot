package main

import (
	"math/rand"
)

// VehicleType keeps track of the different kind of vehicles
// later we can add additional kinds, ex. ElectricType, AccesibleType, etc
type VehicleType int

const (
	MotorcycleType VehicleType = iota
	CarType
	BusType
)

func (vt VehicleType) String() string {
	return [...]string{"Motorcycle", "Car", "Bus"}[vt]
}

//LicensePlate holds a vehilcle license plate number
type LicensePlate string

//Vehicle is the interface for all vehicles
type Vehicle interface {
	Type() VehicleType
	Plate() LicensePlate
	Park(*Lot) (bool, ParkingSpot)
}

//VehicleInfo holds the main information for all vehicles
type VehicleInfo struct {
	LicensePlate
	Type VehicleType
}

//Plate returns a vehicle's license plate number
func (vi VehicleInfo) Plate() LicensePlate {
	return vi.LicensePlate
}

//NewPlate creates a random plate number; obviously check for dupolicate plates...
func NewPlate() LicensePlate {
	var chars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 6)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return LicensePlate(b)
}
