package main

import (
  "bufio"
  "os"
  "strings"
  "strconv"
  "fmt"

  "parkinglot/entity"
  "parkinglot/repository"
  "parkinglot/service"
)

func main() {
  var parkingService *service.ParkingService
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    command := scanner.Text()
    if command == "exit" {
      break
    }
    parts := strings.Fields(command)
    switch parts[0] {
    case "create_parking_lot":
      id := parts[1]
      noOfFloors, _ := strconv.Atoi(parts[2])
      noOfSlotsPerFloor, _ := strconv.Atoi(parts[3])
      repo := repository.NewParkingRepository(id, noOfFloors, noOfSlotsPerFloor)
      parkingService = service.NewParkingService(repo)
      fmt.Printf("Created parking lot with %d floors and %d slots per floor\n", noOfFloors, noOfSlotsPerFloor)
    case "park_vehicle":
      vehicleType := entity.VehicleType(parts[1])
      regNo := parts[2]
      color := parts[3]
      vehicle := &entity.Vehicle{Type: vehicleType, RegNo: regNo, Color: color}
      parkingService.ParkVehicle(vehicle)
    case "unpark_vehicle":
      ticketID := parts[1]
      parkingService.UnParkVehicle(ticketID)
    case "display":
      displayType := parts[1]
      vehicleType := entity.VehicleType(parts[2])
      switch displayType {
      case "free_count":
        parkingService.DisplayFreeCount(vehicleType)
      case "free_slots":
        parkingService.DisplayFreeSlots(vehicleType)
      case "occupied_slots":
        parkingService.DisplayOccupiedSlots(vehicleType)
      default:
        fmt.Println("Invalid display type")
      }
    default:
      fmt.Println("Invalid command")
    }
  }
}
