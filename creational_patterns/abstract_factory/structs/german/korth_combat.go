package german

import "fmt"

type KorthCombat struct {
}

func (m *KorthCombat) SpinCylinder() {
	fmt.Println("Willst du etwa angeben??)")
}

func (m *KorthCombat) Reload() {
	fmt.Println("Viel Glück")
}

func (m *KorthCombat) Fire() {
	fmt.Println("Kopfschuss")
}
