package american

import "fmt"

type TAC50 struct {
}

func (t *TAC50) AdjustScope() {
	fmt.Println("your scope perfectly adjust")
}

func (t *TAC50) Reload() {
	fmt.Println("all 5 bullets in rifle")
}

func (t *TAC50) Fire() {
	fmt.Println("terrorist eleminated")
}
