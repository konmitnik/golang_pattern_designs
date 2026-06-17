package interfaces

type Builder interface {
	StartBuilding()
	InstallCPU(string) error
	InstallRAM(int) error
	InstallGPU(string) error
	InstallROM(int) error
	InstallDrive() error
	ConnectCables() error
	MountIntoCase() error
}
