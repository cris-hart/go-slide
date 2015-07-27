package main

import (
	"fmt"
	"bitbucket.org/tux-eithel/slides/animal"
)

func main() {
	
	dog := &Animal{
		"Fufy"
	}
	
	fmt.Println(dog.MyNameIs())
	
}