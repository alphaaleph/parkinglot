package main

// SpotType keeps track of the different kind of parking spots later we can add additional kinds,
// ex. ChargingSpot for electric vehicles, AccesibleSpot for handicapped person, etc
type SpotType int

const (
	MotorcycleSpot SpotType = iota
	CompactSpot
	LargeSpot
)

func (st SpotType) String() string {
	return [...]string{"Motorcycle", "Compact", "Large"}[st]
}

//ISpot the Spot definition
type ISpot interface {
	IsEmpty() bool
}

//ParkingSpot keeps track of a spot information
type ParkingSpot struct {
	Floor int
	Row   int
	Spot
	Type     SpotType
	Occupied *Vehicle
}

//IsEmpty checks if a ParkingSpot is empty
func (ps ParkingSpot) IsEmpty() bool {
	if ps.Occupied == nil {
		return true
	}
	return false
}
