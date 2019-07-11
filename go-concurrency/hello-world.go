package main

import "fmt"

func main() {

	fmt.Println("Sequential Hello World")


	func() {
		fmt.Println("Hello")
	}()
	func() {
		fmt.Println("World")
	}()
}
