package main

import "fmt"

func main() {

	go (func() {
		fmt.Println("first Goroutine")
	})()

	go (func() {
		fmt.Println("second Goroutine")
	})()

	fmt.Println("Hi!")
}
