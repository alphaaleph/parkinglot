package main

import (
	"fmt"
)

//AvailableSpots displays a list of available spots based on SpotType
func (l Lot) AvailableSpots() {
	lm, lc, ll := len(l.MotorcycleSpots), len(l.CompactSpots), len(l.LargeSpots)

	fmt.Println()
	if lm == 0 && lc == 0 && ll == 0 {
		fmt.Println("FULL LOT")
	} else {
		fmt.Println("-- Number of motorcycle spots available:", lm)
		fmt.Println("-- Number of compact spots available:", lc)
		fmt.Println("-- Number of large spots available:", ll)
	}
	fmt.Println()
}

// UsedSpot prints out the information of a spot that was just taken
func (l Lot) UsedSpot(ps ParkingSpot) {
	vtype := (*ps.Occupied).Type()
	switch vtype {
	case BusType:
		fmt.Printf("%v with plate: %v parked on floor: %d, row: %d, starting at spot: %d for 5 consecutive spaces\n",
			(*ps.Occupied).Type(), (*ps.Occupied).Plate(), ps.Floor, ps.Row, ps.Spot)
	default:
		fmt.Printf("%v with plate: %v parked on floor: %d, row %d, spot: %d\n",
			(*ps.Occupied).Type(), (*ps.Occupied).Plate(), ps.Floor, ps.Row, ps.Spot)
	}
}
