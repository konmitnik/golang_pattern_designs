package structs

type Dog struct {
	Name  string
	Breed string
}

func (d *Dog) Speak() string {
	return "Woof"
}
