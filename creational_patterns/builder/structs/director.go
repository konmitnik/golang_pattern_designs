package structs

import (
	"builder/interfaces"
	"errors"
)

type Director struct {
}

func (d *Director) MakePC(builder interfaces.Builder, cpu string, gpu string, ram int, rom int) error {
	if builder == nil {
		return errors.New("no builder passed to method")
	}

	builder.StartBuilding()

	if err := builder.InstallCPU(cpu); err != nil {
		return err
	}

	if err := builder.InstallRAM(ram); err != nil {
		return err
	}

	if err := builder.InstallGPU(gpu); err != nil {
		return err
	}

	if err := builder.InstallROM(rom); err != nil {
		return err
	}

	if err := builder.InstallDrive(); err != nil {
		return err
	}

	if err := builder.ConnectCables(); err != nil {
		return err
	}

	if err := builder.MountIntoCase(); err != nil {
		return err
	}

	return nil
}
