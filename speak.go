package main

import (
	"fmt"

	"bitbucket.org/tux-eithel/slides/animal"
	"bitbucket.org/tux-eithel/slides/human"
)

func main() {
	dog := &animal.Animal{
		"Fufy",
	}
	cris := &human.Human{
		"Cris",
		"ciao Fiori",
	}
	Print(dog)
	Print(cris)
}

func Print(x interface{}) {
	app := x.(*animal.Animal)
	fmt.Println(app.MyNameIs())
}
