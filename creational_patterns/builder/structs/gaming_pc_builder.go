package structs

import "fmt"

type GamingPCBuilder struct {
	pc GamingPC
}

func (g *GamingPCBuilder) StartBuilding() {
	g.pc = GamingPC{}
}

func (g *GamingPCBuilder) InstallCPU(name string) error {
	g.pc.CPU = name
	fmt.Println(name, "installed")

	return nil
}

func (g *GamingPCBuilder) InstallRAM(capacity int) error {
	g.pc.RAM = capacity
	fmt.Println(capacity, "of RAM installed")

	return nil
}

func (g *GamingPCBuilder) InstallGPU(name string) error {
	g.pc.GPU = name
	fmt.Println(name, "installed")

	return nil
}

func (g *GamingPCBuilder) InstallROM(capacity int) error {
	g.pc.ROM = capacity
	fmt.Println(capacity, "of ROM installed")

	return nil
}

func (g *GamingPCBuilder) InstallDrive() error {
	return nil
}

func (g *GamingPCBuilder) ConnectCables() error {
	g.pc.CablesConnected = true
	fmt.Println("cables connected with power unit")

	return nil
}

func (g *GamingPCBuilder) MountIntoCase() error {
	g.pc.MountedIntoCase = true
	fmt.Println("now its looking cool")

	return nil
}

func (g *GamingPCBuilder) GetPC() GamingPC {
	return g.pc
}
