package german

import "abstract_factory/interfaces"

type Factory struct {
}

func (f *Factory) CreateMachineGun() interfaces.MachineGun {
	return &G36{}
}

func (f *Factory) CreateRevolver() interfaces.Revolver {
	return &KorthCombat{}
}

func (f *Factory) CreateSniperRifle() interfaces.SniperRifle {
	return &SSG3000{}
}
