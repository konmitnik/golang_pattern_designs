package american

import "fmt"

type M4A1 struct {
}

func (g *M4A1) SwitchFireMode() {
	fmt.Println("fire mode switched")
}

func (g *M4A1) Reload() {
	fmt.Println("reloaded")
}

func (g *M4A1) Fire() {
	fmt.Println("terrorist killed")
}
