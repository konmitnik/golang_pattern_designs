package structs

import "fmt"

type OfficePCBuilder struct {
	pc OfficePC
}

func (o *OfficePCBuilder) StartBuilding() {
	o.pc = OfficePC{}
}

func (o *OfficePCBuilder) InstallCPU(name string) error {
	o.pc.CPU = name
	fmt.Println(name, "installed")

	return nil
}

func (o *OfficePCBuilder) InstallRAM(capacity int) error {
	o.pc.RAM = capacity
	fmt.Println(capacity, "of RAM installed")

	return nil
}

func (o *OfficePCBuilder) InstallGPU(_ string) error {
	return nil
}

func (o *OfficePCBuilder) InstallROM(capacity int) error {
	o.pc.ROM = capacity
	fmt.Println(capacity, "of ROM installed")

	return nil
}

func (o *OfficePCBuilder) InstallDrive() error {
	o.pc.DriveInstalled = true
	fmt.Println("drive installed")

	return nil
}

func (o *OfficePCBuilder) ConnectCables() error {
	o.pc.CablesConnected = true
	fmt.Println("cables connected with power unit")

	return nil
}

func (o *OfficePCBuilder) MountIntoCase() error {
	o.pc.MountedIntoCase = true
	fmt.Println("all components mounted into case")

	return nil
}

func (o *OfficePCBuilder) GetPC() OfficePC {
	return o.pc
}
