package model

import (
	"github.com/google/uuid"
)

type ParkingFloor struct {
	id           string
	name         string
	parkingSlots map[ParkingslotType][]ParkingSlot
}

func NewParkingFloor(name string, parkingSlots map[ParkingslotType][]ParkingSlot) ParkingFloor {
	return ParkingFloor{
		id:           uuid.New().String(),
		name:         name,
		parkingSlots: parkingSlots,
	}
}

func (pf ParkingFloor) GetParkingSlots() map[ParkingslotType][]ParkingSlot {
	return pf.parkingSlots
}

func (pf ParkingFloor) GetName() string {
	return pf.name
}

func (pf ParkingFloor) GetFloorID() string {
	return pf.id
}

func (pf *ParkingFloor) SetParkingSlots(slotType ParkingslotType, slots []ParkingSlot) {
	m := make(map[ParkingslotType][]ParkingSlot)
	m[slotType] = slots
	pf.parkingSlots = m
}
