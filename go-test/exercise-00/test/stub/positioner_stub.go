package stub

import (
	"testdoubles/positioner"
)

type PositionerStub struct{}

func NewPositionerStub() *PositionerStub {
	return &PositionerStub{}
}

func (p *PositionerStub) GetLinearDistance(from, to *positioner.Position) (linearDistance float64) {
	linearDistance = 10.00
	return
}
