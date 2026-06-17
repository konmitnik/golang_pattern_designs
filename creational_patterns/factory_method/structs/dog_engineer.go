package structs

import "factory_method/interfaces"

type DogEngineer struct{}

func (d DogEngineer) CreateAnimal(name string, breed string) interfaces.Animal {
	return &Dog{Name: name, Breed: breed}
}
