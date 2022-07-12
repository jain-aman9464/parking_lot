package service

import (
	"fmt"
	"github.com/tokopedia/test/parking_lot/src/parking_lot/model"
	"github.com/tokopedia/test/parking_lot/src/parking_lot/repo"
)

type TicketService struct {
	ticketRepo     repo.TicketRepo
	parkingLotRepo repo.ParkingLotRepo
}

func NewTicketService(lotRepo repo.ParkingLotRepo, ticketRepo repo.TicketRepo) TicketService {
	return TicketService{parkingLotRepo: lotRepo, ticketRepo: ticketRepo}
}

func (t TicketService) AssignTicket(lot *model.ParkingLot, vehicle model.Vehicle, parkingSlotType model.ParkingslotType) model.Ticket {
	slot := t.parkingLotRepo.GetParkingSlotForVehicleAndPark(lot, vehicle, parkingSlotType)
	if slot == nil {
		fmt.Println("Parking full for slot type: ", parkingSlotType)
		return model.Ticket{}
	}

	return t.ticketRepo.AssignTicket(vehicle, *slot)
}

func (t TicketService) ScanAndPay(vehicle model.Vehicle) float64 {
	vehicleFloor := t.parkingLotRepo.GetVehicleFloor(vehicle)
	return t.ticketRepo.ScanAndPay(vehicle, vehicleFloor)
}
