package service

import (
	"github.com/tokopedia/test/parking_lot/src/parking_lot/model"
	"github.com/tokopedia/test/parking_lot/src/parking_lot/repo"
)

type ParkingLotService struct {
	parkLotRepo repo.ParkingLotRepo
}

func NewParkingLotService(lotRepo repo.ParkingLotRepo) ParkingLotService {
	return ParkingLotService{parkLotRepo: lotRepo}
}

func (pf ParkingLotService) PickCorrectSlot(category model.VehicleCategory) model.ParkingslotType {
	if category == model.TwoWheeler {
		return model.TwoWheelerType
	}

	if category == model.HatchBack || category == model.Sedan {
		return model.Compact
	}

	if category == model.SUV {
		return model.Medium
	}

	if category == model.Bus {
		return model.Large
	}

	return -1
}
