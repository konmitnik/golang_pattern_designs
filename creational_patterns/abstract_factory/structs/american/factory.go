package american

import "abstract_factory/interfaces"

type Factory struct {
}

func (f *Factory) CreateMachineGun() interfaces.MachineGun {
	return &M4A1{}
}

func (f *Factory) CreateRevolver() interfaces.Revolver {
	return &MnP9{}
}

func (f *Factory) CreateSniperRifle() interfaces.SniperRifle {
	return &TAC50{}
}
