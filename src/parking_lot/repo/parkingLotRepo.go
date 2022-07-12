package repo

import (
	"github.com/tokopedia/test/parking_lot/src/parking_lot/model"
)

type ParkingLotRepo struct {
	parkingLotDataMap map[string][]model.ParkingFloor
	vehicleFloorMap   map[string]model.ParkingFloor
}

func NewParkingLotRepo() ParkingLotRepo {
	return ParkingLotRepo{
		parkingLotDataMap: make(map[string][]model.ParkingFloor, 0),
		vehicleFloorMap:   make(map[string]model.ParkingFloor, 0),
	}
}

func (p ParkingLotRepo) AddFloors(floor model.ParkingFloor, lot *model.ParkingLot) {
	floors := p.parkingLotDataMap[lot.GetParkingLotID()]
	floors = append(floors, floor)
	p.parkingLotDataMap[lot.GetParkingLotID()] = floors
	lot.SetParkingFloor(floors)
}

func (p ParkingLotRepo) RemoveFloors(floor model.ParkingFloor, lot *model.ParkingLot) {
	floors := p.parkingLotDataMap[lot.GetParkingLotID()]
	for idx, floorVal := range floors {
		if floorVal.GetFloorID() == floor.GetFloorID() {
			floors = append(floors[:idx], floors[idx+1])
		}
	}

	p.parkingLotDataMap[lot.GetName()] = floors
	lot.SetParkingFloor(floors)
}

func (p ParkingLotRepo) GetParkingSlotForVehicleAndPark(lot *model.ParkingLot, vehicle model.Vehicle, parkingSlotType model.ParkingslotType) *model.ParkingSlot {
	var slot *model.ParkingSlot
	floors := lot.GetParkingFloors()
	//floors := p.parkingLotDataMap[lot.GetParkingLotID()]
	for _, floor := range floors {
		relevantParkingSlot := floor.GetParkingSlots()[parkingSlotType]
		for idx, s := range relevantParkingSlot {
			if s.IsAvailable() {
				slot = &s
				slot.AddVehicle(vehicle, floor)
				relevantParkingSlot[idx] = *slot
				floor.SetParkingSlots(parkingSlotType, relevantParkingSlot)
				p.vehicleFloorMap[vehicle.GetVehicleNumber()] = floor
				break
			}
		}
	}

	return slot
}

func (p ParkingLotRepo) GetVehicleFloor(vehicle model.Vehicle) model.ParkingFloor {
	return p.vehicleFloorMap[vehicle.GetVehicleNumber()]
}
