package entity

type ParkingLot struct {
  Id string
  Floors []*Floor
}

func NewParkingLot(id string, noOfFloors, noOfSlotsPerFloor int) *ParkingLot {
  parkingLot := &ParkingLot{Id: id}
  for i := 1; i <= noOfFloors; i++ {
    floor := NewFloor(i, noOfSlotsPerFloor)
    parkingLot.Floors = append(parkingLot.Floors, floor)
  }
  return parkingLot
}
