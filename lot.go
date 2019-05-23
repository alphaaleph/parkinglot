package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Spot represent a single parking space in a lot
type Spot int

// Row has a group of spots
type Row struct {
	Spots []Spot
}

// Floor has a group of rows
type Floor struct {
	Rows []Row
}

//ILot the Lot definition
type ILot interface {
	FillSpot(Vehicle)
	Leave()
	AvailableSpots()
	UsedSpots(ParkingSpot)
	FindMotorcycleSpot(int) (bool, ParkingSpot)
	FindCompactSpot(int) (bool, ParkingSpot)
	FindLargeSpot(int) (bool, ParkingSpot)
	FindConsecutiveLargeSpot(int) (bool, ParkingSpot)
}

// Lot keeps track of all the floors, rows, and spots in a parking lot
type Lot struct {
	Floors []Floor
	//keeps track of parked vehicles
	ParkedVehicles map[LicensePlate]ParkingSpot
	//keeping a list of vehicle types for faster search of an available parking spot
	MotorcycleSpots []ParkingSpot
	CompactSpots    []ParkingSpot
	LargeSpots      []ParkingSpot
}

// FillSpot is used to park a vehicle in a random spot
func (l *Lot) FillSpot(v Vehicle) {

	if ok, ps := v.Park(l); ok {
		//spot tracks vehicle
		ps.Occupied = &v

		//track vehicle and spot in map
		l.ParkedVehicles[(*ps.Occupied).Plate()] = ps
		l.UsedSpot(ps)
		ok = false
	}
}

// Leave keeps track of vehicles that left the lot
func (l *Lot) Leave() {
	if len(l.ParkedVehicles) <= 0 {
		fmt.Println("EMPTY LOT")
	}

	//pick random vehicle to leave
	var releaseSpace int
	if len(l.ParkedVehicles) == 1 {
		releaseSpace = 0
	} else {
		rand.Seed(time.Now().UnixNano())
		releaseSpace = rand.Intn(len(l.ParkedVehicles) - 1)
	}

	var plate LicensePlate
	for plate = range l.ParkedVehicles {
		if releaseSpace == 0 {
			break
		}
		releaseSpace--
	}

	ps := l.ParkedVehicles[plate]
	delete(l.ParkedVehicles, plate)

	switch ps.Type {
	case MotorcycleSpot:
		l.MotorcycleSpots = append(l.MotorcycleSpots, ps)
	case CompactSpot:
		l.CompactSpots = append(l.CompactSpots, ps)
	case LargeSpot:
		l.LargeSpots = append(l.LargeSpots, ps)
	default:
	}
}

// random space is used to park a vehicle in a randon parking spot
func (l Lot) randomSpace(spotLen int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(spotLen - 1)
}
