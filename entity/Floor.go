package entity

type Floor struct {
  FloorNo int
  Slots []*Slot
}

func NewFloor(floorNo, noOfSlots int) *Floor {
  floor := &Floor{FloorNo: floorNo}
  for j := 1; j <= noOfSlots; j++ {
    var slotType VehicleType
    if j == 1 {
        slotType = Truck
    } else if j == 2 || j == 3 {
        slotType = Bike
    } else {
        slotType = Car
    }
    slot := &Slot{SlotNo: j, IsOccupied: false, SlotType: slotType}
    floor.Slots = append(floor.Slots, slot)
  }
  return floor
}

func (this *Floor) AssignVacantSlot(vehicle *Vehicle) *Slot {
  for _, slot := range this.Slots {
    if slot.SlotType == vehicle.Type && slot.IsOccupied == false {
      slot.Vehicle = vehicle
      slot.IsOccupied = true
      return slot
    }
  }
  return nil
}

func (this *Floor) GetFreeSlotCount(vehicleType VehicleType) int {
  count := 0
  for _, slot := range this.Slots {
    if slot.SlotType == vehicleType && slot.IsOccupied == false {
      count++
    }
  }
  return count
}

func (this *Floor) GetFreeSlots(vehicleType VehicleType) []int {
  result := []int{}
  for _, slot := range this.Slots {
    if slot.SlotType == vehicleType && slot.IsOccupied == false {
      result = append(result, slot.SlotNo)
    }
  }
  return result
}

func (this *Floor) GetOccupiedSlots(vehicleType VehicleType) []int {
  result := []int{}
  for _, slot := range this.Slots {
    if slot.SlotType == vehicleType && slot.IsOccupied == true {
      result = append(result, slot.SlotNo)
    }
  }
  return result
}
