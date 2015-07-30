package main

import (
	"fmt"

	"bitbucket.org/tux-eithel/slides/animal"
	"bitbucket.org/tux-eithel/slides/human"
)

type ispeak interface {
	MyNameIs() string
	Speak() string
}

func main() {
	dog := &animal.Animal{
		"Fufy",
	}
	cris := &human.Human{
		"Cris",
		"ciao devs",
	}
	bird := &animal.Parrot{
		animal.Animal{"Kirbi"},
		"hi hi",
	}

	Print(dog)
	Print(cris)
	Print(bird)

	PrintI(dog)
}

func Print(x interface{}) {

	switch t := x.(type) {

	case *human.Human:
		fmt.Println(t.MyNameIs())
		fmt.Println(t.Speak())

	case *animal.Parrot:
		fmt.Println(t.MyNameIs())
		fmt.Println(t.Speak())

	default:
		panic("not supported")

	}

}

func PrintI(t ispeak) {
	fmt.Println(t.MyNameIs())
	fmt.Println(t.Speak())
}
