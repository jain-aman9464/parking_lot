package model

import (
	"github.com/google/uuid"
	"time"
)

type Ticket struct {
	ticketNumber string
	startTime    time.Time
	endTime      time.Time
	vehicle      Vehicle
	parkingSlot  ParkingSlot
}

func NewTicket(vehicle Vehicle, slot ParkingSlot) Ticket {
	timeNow := time.Now()
	return Ticket{
		ticketNumber: vehicle.GetVehicleNumber() + uuid.New().String(),
		startTime:    timeNow,
		vehicle:      vehicle,
		parkingSlot:  slot,
	}
}

func (t Ticket) GetStartTime() time.Time {
	return t.startTime
}

func (t Ticket) GetEndTime() time.Time {
	return t.endTime
}

func (t Ticket) GetParkingSlot() ParkingSlot {
	return t.parkingSlot
}

func (t *Ticket) SetEndTime(ti time.Time) {
	t.endTime = ti
}

func (t Ticket) GetTicketNum() string {
	return t.ticketNumber
}

func (t Ticket) GetVehicle() Vehicle {
	return t.vehicle
}
