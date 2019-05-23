package main

import "sort"

// FindMotorcycleSpot will look for an available motorcycle parking spot in the lot
func (l *Lot) FindMotorcycleSpot(neededSpaces int) (bool, ParkingSpot) {
	//no spots left
	if len(l.MotorcycleSpots) <= 0 {
		return false, ParkingSpot{}
	}

	//park in random location
	space := l.randomSpace(len(l.MotorcycleSpots))
	ps := l.MotorcycleSpots[space]

	//remove from lot list
	l.MotorcycleSpots = append(l.MotorcycleSpots[:space], l.MotorcycleSpots[space+neededSpaces:]...)
	return true, ps
}

// FindCompactSpot will look for an available compact parking spot in the lot
func (l *Lot) FindCompactSpot(neededSpaces int) (bool, ParkingSpot) {
	//no spots left
	if len(l.CompactSpots) <= 0 {
		return false, ParkingSpot{}
	}

	//park in random location
	var space int
	if len(l.CompactSpots) == 1 {
		space = 0
	} else {
		space = l.randomSpace(len(l.CompactSpots))
	}
	ps := l.CompactSpots[space]

	//remove from lot list
	l.CompactSpots = append(l.CompactSpots[:space], l.CompactSpots[space+neededSpaces:]...)
	return true, ps
}

// FindLargeSpot will look for an available large parking spot in the lot
func (l *Lot) FindLargeSpot(neededSpaces int) (bool, ParkingSpot) {
	//no spots left
	if len(l.LargeSpots) <= 0 {
		return false, ParkingSpot{}
	}

	//park in random location
	space := l.randomSpace(len(l.LargeSpots))
	ps := l.LargeSpots[space]

	//remove from lot list
	l.LargeSpots = append(l.LargeSpots[:space], l.LargeSpots[space+neededSpaces:]...)
	return true, ps
}

// FindConsecutiveLargeSpot will look for parking spots that are consecutively available
func (l *Lot) FindConsecutiveLargeSpot(neededSpaces int) (bool, ParkingSpot) {
	//no enough spaces for large vehicle
	if len(l.LargeSpots) <= neededSpaces-1 {
		return false, ParkingSpot{}
	}

	//check if there are enough consecutive spaces
	sort.Slice(l.LargeSpots[:], func(i, j int) bool {
		return l.LargeSpots[i].Spot < l.LargeSpots[j].Spot
	})

	var hspot int
	if hspot = l.headSpot(neededSpaces); hspot < 0 {
		return false, ParkingSpot{}
	}

	//park in starting consecutive spots
	ps := l.LargeSpots[hspot]
	//remove from list all spaces
	l.LargeSpots = append(l.LargeSpots[:hspot], l.LargeSpots[hspot+neededSpaces:]...)
	return true, ps
}

// headSpot will return the head parking spot of a group of gonsecutively used parking spots
func (l *Lot) headSpot(neededSpaces int) int {
	pos := 0
	count := 1
	head := l.LargeSpots[pos]

	for i := 1; i < len(l.LargeSpots)-neededSpaces; i++ {
		if (head.Spot + Spot(i)) == l.LargeSpots[i].Spot {
			count++
		} else {
			head = l.LargeSpots[i]
			pos = 1
		}
		if count == neededSpaces {
			return pos
		}
	}
	return -1
}
