package main

import (
	"fmt"
)

func main() {

	fmt.Println("Concurrent Hello World")

	ch := make(chan bool)
	lch := make(chan bool)

	// note the go call in front of function
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Hello")
		}
		lch <- true
	}()
	go func() {
		<- lch
		for i := 0; i < 10; i++ {
			fmt.Println("World")
		}
		ch <- true
	}()

	fmt.Println("before finish 1")
	<-ch
	fmt.Println("before finish 2")

}
