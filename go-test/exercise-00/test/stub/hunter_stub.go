package stub

import "testdoubles/prey"

type HunterStub struct {
}

func NewHunterStub() *HunterStub {
	return &HunterStub{}
}

func (p *HunterStub) Hunt(prey prey.Prey) (err error) {
	return
}
