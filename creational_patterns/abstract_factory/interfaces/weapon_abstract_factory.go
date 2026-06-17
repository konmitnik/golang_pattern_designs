package interfaces

type WeaponAbstractFactory interface {
	CreateMachineGun() MachineGun
	CreateRevolver() Revolver
	CreateSniperRifle() SniperRifle
}
