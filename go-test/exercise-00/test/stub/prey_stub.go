package stub

import (
	"testdoubles/positioner"
)

type PreyStub struct {
	speed    float64
	position *positioner.Position
}

func NewPreyStub() *PreyStub {
	return &PreyStub{}
}

func (p *PreyStub) GetSpeed() (speed float64) {
	speed = p.speed
	return
}

func (p *PreyStub) GetPosition() (position *positioner.Position) {
	position = p.position
	return
}
