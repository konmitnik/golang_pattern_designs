package main

import (
	"abstract_factory/interfaces"
	"abstract_factory/structs/russian"
)

func main() {
	var factory interfaces.WeaponAbstractFactory = &russian.Factory{}

	var machineGun interfaces.MachineGun = factory.CreateMachineGun()
	var revolver interfaces.Revolver = factory.CreateRevolver()
	var sniperRifle interfaces.SniperRifle = factory.CreateSniperRifle()

	machineGun.SwitchFireMode()
	revolver.SpinCylinder()
	sniperRifle.AdjustScope()
}
