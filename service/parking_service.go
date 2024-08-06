package service

import (
  "fmt"
  "parkinglot/repository"
  "parkinglot/entity"
)

type ParkingService struct {
  Repo *repository.ParkingRepository
}

func NewParkingService(repo *repository.ParkingRepository) *ParkingService {
  return &ParkingService{Repo: repo}
}

func (service *ParkingService) ParkVehicle(vehicle *entity.Vehicle) {
  ticketId := service.Repo.ParkVehicle(vehicle)
  if ticketId == "" {
    fmt.Println("Parking lot is full")
  } else {
    fmt.Println("Parked Vehicle. Ticket Id: %s\n", ticketId)
  }
}

func (service *ParkingService) UnParkVehicle(ticketId string) {
  vehicle, isValidTicket := service.Repo.UnParkVehicle(ticketId)
  if !isValidTicket {
    fmt.Println("Invalid Ticket")
  } else {
    fmt.Println("UnParked vehicle with registration number: %s and color: %s\n", vehicle.RegNo, vehicle.Color)
  }
}

func (service *ParkingService) DisplayFreeCount(vehicleType entity.VehicleType) {
  freeSlotsCount := service.Repo.GetFreeSlotCount(vehicleType)
  for floorNo, count := range freeSlotsCount {
    fmt.Printf("No. of free slots for %s on Floor %d: %d\n", vehicleType, floorNo, count)
  }
}

func (service *ParkingService) DisplayFreeSlots(vehicleType entity.VehicleType) {
	freeSlots := service.Repo.GetFreeSlots(vehicleType)
	for floorNo, slots := range freeSlots {
		fmt.Printf("Free slots for %s on Floor %d: %v\n", vehicleType, floorNo, slots)
	}
}

func (service *ParkingService) DisplayOccupiedSlots(vehicleType entity.VehicleType) {
	occupiedSlots := service.Repo.GetOccupiedSlots(vehicleType)
	for floorNo, slots := range occupiedSlots {
		fmt.Printf("Occupied slots for %s on Floor %d: %v\n", vehicleType, floorNo, slots)
	}
}
