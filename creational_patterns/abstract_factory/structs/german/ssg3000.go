package german

import "fmt"

type SSG3000 struct {
}

func (t *SSG3000) AdjustScope() {
	fmt.Println("Ihr Visier ist eingestellt")
}

func (t *SSG3000) Reload() {
	fmt.Println("Das Magazin wird in das Gewehr eingesetzt.")
}

func (t *SSG3000) Fire() {
	fmt.Println("Das Ziel wird getroffen")
}
