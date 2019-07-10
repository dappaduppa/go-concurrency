package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	//var li list.List
	//li = new()

	var sl []int = make([]int, 10, 20)
	for i := range sl {
		sl[i] = 100 - i*2
	}

	fmt.Println("before=", sl)

	intc := make(chan int)
	go func() {
		fmt.Println("pretending it is heavy sort function...5 seconds sleeping...")
		time.Sleep(time.Second * time.Duration(5))
		sort.Slice(sl, func(i, j int) bool { return sl[i] < sl[j] })
		intc <- 1
	}()

	go func() {
		for {
			time.Sleep(time.Second * 1)
			fmt.Print(".")
		}
	}()

	<-intc

	fmt.Println()
	fmt.Println("after=", sl)

}
