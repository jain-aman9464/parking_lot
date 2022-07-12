package model

type VehicleCategory int

const (
	TwoWheeler VehicleCategory = iota
	HatchBack
	Sedan
	SUV
	Bus
)

type Vehicle struct {
	vehicleNum      string
	vehicleCategory VehicleCategory
}

func NewVehicle(num string) Vehicle {
	return Vehicle{
		vehicleNum: num,
	}
}

func (v *Vehicle) SetVehicleNum(num string) {
	v.vehicleNum = num
}

func (v *Vehicle) SetVehicleCategory(category VehicleCategory) {
	v.vehicleCategory = category
}

func (v Vehicle) GetVehicleCategory() VehicleCategory {
	return v.vehicleCategory
}

func (v Vehicle) GetVehicleNumber() string {
	return v.vehicleNum
}
