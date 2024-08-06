package repository

import (
  "fmt"
  "strings"
  "strconv"
  "parkinglot/entity"
)

type ParkingRepository struct {
  ParkingLot *entity.ParkingLot
}

func NewParkingRepository(id string, noOfFloors, noOfSlotsPerFloor int) *ParkingRepository {
  return &ParkingRepository{ParkingLot: entity.NewParkingLot(id, noOfFloors, noOfSlotsPerFloor)}
}

func (repo *ParkingRepository) ParkVehicle(vehicle *entity.Vehicle) string {
  for _, floor := range repo.ParkingLot.Floors {
    slot := floor.AssignVacantSlot(vehicle)
    if slot != nil {
      ticketId := fmt.Sprintf("%s_%d_%d", repo.ParkingLot.Id, floor.FloorNo, slot.SlotNo)
      return ticketId
    }
  }
  return ""
}

func (repo *ParkingRepository) UnParkVehicle(ticketId string) (*entity.Vehicle, bool) {
  // Get information from ticket id
  ticketParts := strings.Split(ticketId, "_")
  if len(ticketParts) != 3 {
    return nil, false
  }

  // Get floor and slot info
  floorNo, _ := strconv.Atoi(ticketParts[1])
  slotNo, _ := strconv.Atoi(ticketParts[2])
  floor := repo.ParkingLot.Floors[floorNo - 1]
  slot := floor.Slots[slotNo - 1]

  // Handle the case when slot is unoccupied
  if slot.IsOccupied == false {
    return nil, false
  }

  // UnPark the vehicle
  vehicle := slot.Vehicle
  slot.Vehicle = nil
  slot.IsOccupied = false
  return vehicle, true
}

func (repo *ParkingRepository) GetFreeSlotCount(vehicleType entity.VehicleType) map[int]int {
  result := make(map[int]int)
  for _, floor := range repo.ParkingLot.Floors {
    slotCount := floor.GetFreeSlotCount(vehicleType)
    if slotCount != 0 {
      result[floor.FloorNo] = slotCount
    }
  }
  return result
}

func (repo *ParkingRepository) GetFreeSlots(vehicleType entity.VehicleType) map[int][]int {
  result := make(map[int][]int)
  for _, floor := range repo.ParkingLot.Floors {
    freeSlots := floor.GetFreeSlots(vehicleType)
    if len(freeSlots) > 0 {
      result[floor.FloorNo] = freeSlots
    }
  }
  return result
}

func (repo *ParkingRepository) GetOccupiedSlots(vehicleType entity.VehicleType) map[int][]int {
  result := make(map[int][]int)
  for _, floor := range repo.ParkingLot.Floors {
    occupiedSlots := floor.GetOccupiedSlots(vehicleType)
    if len(occupiedSlots) > 0 {
      result[floor.FloorNo] = occupiedSlots
    }
  }
  return result
}
