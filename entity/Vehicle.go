package entity

type VehicleType string

const (
  Car   VehicleType = "CAR"
  Bike  VehicleType = "Bike"
  Truck VehicleType = "Truck"
)

type Vehicle struct {
  Type  VehicleType
  RegNo string
  Color string
}
