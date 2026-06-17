package russian

import "fmt"

type RSH12 struct {
}

func (r *RSH12) SpinCylinder() {
	fmt.Println("are you flexing?)")
}

func (r *RSH12) Reload() {
	fmt.Println("very lony reload complete")
}

func (r *RSH12) Fire() {
	fmt.Println("powerful shot")
}
