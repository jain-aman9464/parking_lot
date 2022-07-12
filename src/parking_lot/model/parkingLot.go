package model

import (
	"github.com/google/uuid"
)

type ParkingLot struct {
	id               string
	nameOfParkingLot string
	parkingFloors    []ParkingFloor
}

func NewParkingLot(nameOfParkingLot string, parkingFloors []ParkingFloor) ParkingLot {
	return ParkingLot{
		id:               uuid.New().String(),
		nameOfParkingLot: nameOfParkingLot,
		parkingFloors:    parkingFloors,
	}
}

func (p ParkingLot) GetParkingFloors() []ParkingFloor {
	return p.parkingFloors
}

func (p ParkingLot) GetParkingLotID() string {
	return p.id
}

func (p ParkingLot) GetName() string {
	return p.nameOfParkingLot
}

func (p *ParkingLot) SetParkingFloor(floors []ParkingFloor) {
	p.parkingFloors = floors
}
