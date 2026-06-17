package structs

type Cat struct {
	Name  string
	Breed string
}

func (c *Cat) Speak() string {
	return "Meow"
}
