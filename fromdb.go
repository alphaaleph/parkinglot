package main

// BuildLot example with lot which consists of 2 floors.
//		First floor has 2 rows:
//			row 1 with 12 motorcycles spaces	: numbered 101-112
//			row 2 with 10 compact spaces		: numbered 113-122
//			row 3 with 5 large spaces			: numbered 123-127
//		Second floor has 2 rows:
//			row 1 with 10 compact spaces						: numbered 201-210
//			row 2 with 4 motorcycle spaces, and 10 large spaces	: numbered 211-214, and 215-224 respectively
//
//	we can read this information from a DB
//
// 	we have a total of:
//		16 motorcycle spaces
//		20 compact spaces
//		15 large spaces
func BuildLot(lot *Lot) {

	//init the parked vehicles
	lot.ParkedVehicles = make(map[LicensePlate]ParkingSpot)

	//add floors
	for f := 1; f < 3; f++ {
		lot.Floors = append(lot.Floors, Floor{})

		switch f {
		case 1:
			// add rows
			for r := 1; r < 4; r++ {
				switch r {
				case 1:
					for m := Spot(101); m < 113; m++ {
						mspace := ParkingSpot{f, r, m, MotorcycleSpot, nil}
						lot.MotorcycleSpots = append(lot.MotorcycleSpots, mspace)
					}
				case 2:
					for c := Spot(113); c < 123; c++ {
						cspace := ParkingSpot{f, r, c, CompactSpot, nil}
						lot.CompactSpots = append(lot.CompactSpots, cspace)
					}
				case 3:
					for l := Spot(123); l < 128; l++ {
						lspace := ParkingSpot{f, r, l, LargeSpot, nil}
						lot.LargeSpots = append(lot.LargeSpots, lspace)
					}
				default:
				}
			}
		case 2:
			// add rows
			for r := 1; r < 3; r++ {
				switch r {
				case 1:
					for c := Spot(201); c < 211; c++ {
						cspace := ParkingSpot{f, r, c, CompactSpot, nil}
						lot.CompactSpots = append(lot.CompactSpots, cspace)
					}
				case 2:
					for m := Spot(211); m < 215; m++ {
						mspace := ParkingSpot{f, r, m, MotorcycleSpot, nil}
						lot.MotorcycleSpots = append(lot.MotorcycleSpots, mspace)
					}
					for l := Spot(215); l < 225; l++ {
						lspace := ParkingSpot{f, r, l, LargeSpot, nil}
						lot.LargeSpots = append(lot.LargeSpots, lspace)
					}
				default:
				}
			}
		default:
		}
	}
}
