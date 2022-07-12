package model

type ParkingSlot struct {
	name            string
	isAvailable     bool
	vehicle         Vehicle
	parkingSlotType ParkingslotType
}

func NewParkingSlot(name string, parkingslotType ParkingslotType) ParkingSlot {
	return ParkingSlot{
		name:            name,
		parkingSlotType: parkingslotType,
		isAvailable:     true,
	}
}

func (p *ParkingSlot) AddVehicle(vehicle Vehicle, floor ParkingFloor) {
	p.vehicle = vehicle
	p.isAvailable = false
}

func (p *ParkingSlot) RemoveVehicle() {
	p.vehicle = Vehicle{}
	p.isAvailable = true
}

func (p ParkingSlot) IsAvailable() bool {
	return p.isAvailable
}

func (p ParkingSlot) GetParkingSlotType() ParkingslotType {
	return p.parkingSlotType
}

func (p ParkingSlot) GetVehicle() Vehicle {
	return p.vehicle
}
