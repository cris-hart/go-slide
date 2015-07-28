package main

import (
	"fmt"
	"sync"
)

func main() {

	snc := &sync.WaitGroup{}
	snc.Add(2)

	go (func(s *sync.WaitGroup) {
		defer s.Done()
		fmt.Println("first Goroutine")
	})(snc)

	go (func(s *sync.WaitGroup) {
		defer s.Done()
		fmt.Println("second Goroutine")
	})(snc)

	snc.Wait()
	fmt.Println("Hi!")
}
