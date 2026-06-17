package russian

import "abstract_factory/interfaces"

type Factory struct {
}

func (f *Factory) CreateMachineGun() interfaces.MachineGun {
	return &AK200{}
}

func (f *Factory) CreateRevolver() interfaces.Revolver {
	return &RSH12{}
}

func (f *Factory) CreateSniperRifle() interfaces.SniperRifle {
	return &SVDK{}
}
