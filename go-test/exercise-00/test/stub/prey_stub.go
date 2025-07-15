package stub

import (
	"testdoubles/positioner"
)

type PreyStub struct {
}

func NewPreyStub() *PreyStub {
	return &PreyStub{}
}

func (t *PreyStub) GetSpeed() (speed float64) {
	speed = 50.00
	return
}

func (t *PreyStub) GetPosition() (position *positioner.Position) {
	position = &positioner.Position{
		X: 10,
		Y: 10,
		Z: 10,
	}
	return
}
