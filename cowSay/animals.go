package main

import (
	"fmt"
)

type Animal struct {
	Name       string
	Food       string
	Locomotion string
	Sound      string
}

func main() {
	animals := []Animal{
		{Name: "cow", Food: "grass", Locomotion: "walk", Sound: "moo"},
		{Name: "bird", Food: "worms", Locomotion: "fly", Sound: "peep"},
		{Name: "snake", Food: "mice", Locomotion: "slither", Sound: "hsss"},
	}

	for {
		fmt.Print("Enter animal name (cow, bird, snake): ")
		var animalName string
		fmt.Scan(&animalName)
		fmt.Print("and information type (eat, move, speak): ")
		var infoType string
		fmt.Scan(&infoType)

		animal := getAnimal(animals, animalName)
		if animal == nil {
			fmt.Println("Invalid animal name. Please try again.")
			continue
		}

		switch infoType {
		case "eat":
			fmt.Println(animal.Food)
		case "move":
			fmt.Println(animal.Locomotion)
		case "speak":
			fmt.Println(animal.Sound)
		default:
			fmt.Println("Invalid information type. Please try again.")
		}
	}
}

func getAnimal(animals []Animal, name string) *Animal {
	for _, animal := range animals {
		if animal.Name == name {
			return &animal
		}
	}
	return nil
}
