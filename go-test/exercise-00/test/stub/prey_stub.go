package stub

import (
	"testdoubles/positioner"
)

type PreyStub struct {
}

func NewPreyStub() *PreyStub {
	return &PreyStub{}
}

func (p *PreyStub) GetSpeed() (speed float64) {
	speed = 10.00
	return
}

func (p *PreyStub) GetPosition() (position *positioner.Position) {
	position = &positioner.Position{}
	return
}
