package american

import "fmt"

type MnP9 struct {
}

func (m *MnP9) SpinCylinder() {
	fmt.Println("are you flexing too?)")
}

func (m *MnP9) Reload() {
	fmt.Println("good luck")
}

func (m *MnP9) Fire() {
	fmt.Println("headshot")
}
