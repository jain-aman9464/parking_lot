package model

type ParkingslotType int

const (
	TwoWheelerType ParkingslotType = iota
	Compact
	Medium
	Large
)

var (
	ParkingCapacityPerSlotType = map[ParkingslotType]int{
		TwoWheelerType: 4,
		Compact:        3,
		Medium:         5,
		Large:          2,
	}

	RateCardPerType = map[ParkingslotType]map[float64]float64{
		TwoWheelerType: {
			2: 10,
			4: 20,
			6: 30,
			8: 40,
		},
		Compact: {
			2: 20,
			4: 40,
			6: 60,
			8: 80,
		},
		Medium: {
			2: 30,
			4: 60,
			6: 90,
			8: 120,
		},
		Large: {
			2: 40,
			4: 80,
			6: 120,
			8: 160,
		},
	}
)

func (v ParkingslotType) GetPriceForParking(duration float64) float64 {
	rateCard := RateCardPerType[v]

	switch v {
	case TwoWheelerType:
		for t, p := range rateCard {
			if t > duration {
				return p
			}
		}

	case Compact:
		for t, p := range rateCard {
			if t > duration {
				return p
			}
		}

	case Medium:
		for t, p := range rateCard {
			if t > duration {
				return p
			}
		}

	case Large:
		for t, p := range rateCard {
			if t > duration {
				return p
			}
		}
	}

	return -1
}
