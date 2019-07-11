package main

import "fmt"

func main() {

	fmt.Println("Concurrent Hello World")

	// note the go call in front of function
	go func() {
		fmt.Println("Hello")
	}()
	go func() {
		fmt.Println("World")
	}()
}
