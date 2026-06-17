package russian

import "fmt"

type AK200 struct {
}

func (a *AK200) SwitchFireMode() {
	fmt.Println("fire mode switched")
}

func (a *AK200) Reload() {
	fmt.Println("reloaded")
}

func (a *AK200) Fire() {
	fmt.Println("puff")
}
