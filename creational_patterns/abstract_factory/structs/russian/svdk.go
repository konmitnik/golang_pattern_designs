package russian

import "fmt"

type SVDK struct {
}

func (s *SVDK) AdjustScope() {
	fmt.Println("good luch with aiming")
}

func (s *SVDK) Reload() {
	fmt.Println("bullet in your rifle")
}

func (s *SVDK) Fire() {
	fmt.Println("360 no scope headshot")
}
