package repo

import (
	"github.com/tokopedia/test/parking_lot/src/parking_lot/model"
	"math"
	"time"
)

type TicketRepo struct {
	vehicleTicketMap map[string]model.Ticket
}

func NewTicketRepo() TicketRepo {
	return TicketRepo{
		vehicleTicketMap: make(map[string]model.Ticket, 0),
	}
}

func (p TicketRepo) AssignTicket(vehicle model.Vehicle, slot model.ParkingSlot) model.Ticket {
	newTicket := model.NewTicket(vehicle, slot)
	p.vehicleTicketMap[vehicle.GetVehicleNumber()] = newTicket

	return newTicket
}

func (p TicketRepo) ScanAndPay(vehicle model.Vehicle, floor model.ParkingFloor) float64 {
	ticket := p.vehicleTicketMap[vehicle.GetVehicleNumber()]
	ticket.SetEndTime(time.Now())
	slots := floor.GetParkingSlots()[ticket.GetParkingSlot().GetParkingSlotType()]
	for idx, slot := range slots {
		if slot.GetVehicle().GetVehicleNumber() == ticket.GetVehicle().GetVehicleNumber() {
			slots = append(slots[:idx], slots[idx+1])
			slot.RemoveVehicle()
			break
		}
	}
	floor.SetParkingSlots(ticket.GetParkingSlot().GetParkingSlotType(), slots)
	duration := math.Ceil(float64(ticket.GetEndTime().Second() - ticket.GetStartTime().Second()))
	price := ticket.GetParkingSlot().GetParkingSlotType().GetPriceForParking(duration)

	return price
}
