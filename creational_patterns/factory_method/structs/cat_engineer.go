package structs

import "factory_method/interfaces"

type CatEngineer struct{}

func (c CatEngineer) CreateAnimal(name string, breed string) interfaces.Animal {
	return &Cat{Name: name, Breed: breed}
}
