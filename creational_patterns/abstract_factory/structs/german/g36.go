package german

import "fmt"

type G36 struct {
}

func (g *G36) SwitchFireMode() {
	fmt.Println("Der Feuermodus wurde geändert")
}

func (g *G36) Reload() {
	fmt.Println("Der Verkaufsautomatenladen wurde geändert")
}

func (g *G36) Fire() {
	fmt.Println("Rat-at-at-at")
}
