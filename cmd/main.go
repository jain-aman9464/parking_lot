package main

import (
	"fmt"
	"github.com/tokopedia/test/parking_lot/src/parking_lot/model"
	"github.com/tokopedia/test/parking_lot/src/parking_lot/repo"
	"github.com/tokopedia/test/parking_lot/src/parking_lot/service"
)

func main() {
	fmt.Println("Start feeding data")
	nameOfParkingLot := "Pintoss parking lot"
	allSlots := make(map[model.ParkingslotType][]model.ParkingSlot)

	compactSlot := make([]model.ParkingSlot, 0)
	c1 := model.NewParkingSlot("c1", model.Compact)
	c2 := model.NewParkingSlot("c2", model.Compact)
	c3 := model.NewParkingSlot("c3", model.Compact)

	compactSlot = append(compactSlot, c1, c2, c3)

	largeSlot := make([]model.ParkingSlot, 0)
	l1 := model.NewParkingSlot("l1", model.Large)
	l2 := model.NewParkingSlot("l2", model.Large)
	l3 := model.NewParkingSlot("l3", model.Large)

	largeSlot = append(largeSlot, l1, l2, l3)

	allSlots[model.Compact] = compactSlot
	allSlots[model.Large] = largeSlot
	parkingFloor := model.NewParkingFloor("1", allSlots)

	parkingFloors := make([]model.ParkingFloor, 0)
	parkingFloors = append(parkingFloors, parkingFloor)

	parkingLot := model.NewParkingLot(nameOfParkingLot, parkingFloors)

	vehicle := model.NewVehicle("KA-01-MA-9999")
	vehicle.SetVehicleCategory(model.HatchBack)

	ticketRepo := repo.NewTicketRepo()
	parkingLotRepo := repo.NewParkingLotRepo()

	parkingLotService := service.NewParkingLotService(parkingLotRepo)
	ticketService := service.NewTicketService(parkingLotRepo, ticketRepo)

	ticket := ticketService.AssignTicket(&parkingLot, vehicle, parkingLotService.PickCorrectSlot(vehicle.GetVehicleCategory()))
	fmt.Println("Ticket num: ", ticket.GetTicketNum(), ticket.GetStartTime(), ticket.GetEndTime())
	//time.Sleep(5 * time.Second)
	//price := ticketService.ScanAndPay(vehicle)
	//fmt.Println("Price: ", price)

	vehicle2 := model.NewVehicle("BMW-9999")
	vehicle2.SetVehicleCategory(model.HatchBack)
	ticket2 := ticketService.AssignTicket(&parkingLot, vehicle2, parkingLotService.PickCorrectSlot(vehicle2.GetVehicleCategory()))
	fmt.Println("Ticket num: ", ticket2.GetTicketNum())

	//time.Sleep(5 * time.Second)
	//price2 := ticketService.ScanAndPay(vehicle2)
	//fmt.Println("Price: ", price2)

	vehicle3 := model.NewVehicle("MERC-9999")
	vehicle3.SetVehicleCategory(model.HatchBack)
	ticket3 := ticketService.AssignTicket(&parkingLot, vehicle3, parkingLotService.PickCorrectSlot(vehicle3.GetVehicleCategory()))
	fmt.Println("Ticket num: ", ticket3.GetTicketNum())

	//time.Sleep(5 * time.Second)
	//price3 := ticketService.ScanAndPay(vehicle3)
	//fmt.Println("Price: ", price3)

	vehicle4 := model.NewVehicle("AUDI-9999")
	vehicle4.SetVehicleCategory(model.HatchBack)
	ticket4 := ticketService.AssignTicket(&parkingLot, vehicle4, parkingLotService.PickCorrectSlot(vehicle4.GetVehicleCategory()))
	fmt.Println("Ticket num: ", ticket4.GetTicketNum())

	//time.Sleep(5 * time.Second)
	//price4 := ticketService.ScanAndPay(vehicle4)
	//fmt.Println("Price: ", price4)

}
