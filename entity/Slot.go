package entity

type Slot struct {
  SlotNo      int
  Vehicle     *Vehicle
  SlotType    VehicleType
  IsOccupied  bool
}
