package main

import (
	"factory_method/interfaces"
	"factory_method/structs"
	"fmt"
)

func main() {
	var catBreeder interfaces.ZooEngineer = structs.CatEngineer{}
	newCat := catBreeder.CreateAnimal("Alice", "Siamese")
	fmt.Printf("cat said %s\n", newCat.Speak())

	var dogBreeder interfaces.ZooEngineer = structs.DogEngineer{}
	newDog := dogBreeder.CreateAnimal("Bobik", "German shepherd")
	fmt.Printf("dog said %s\n", newDog.Speak())
}
