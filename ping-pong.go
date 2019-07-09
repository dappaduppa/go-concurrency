package main

import (
	"fmt"
	"time"
)

func ping(chn chan time.Time) {
	for {
		chn <- time.Now()
		time.Sleep(time.Second * time.Duration(2))
	}
}

func pong(chn chan time.Time) {
	for {
		fmt.Println(<-chn)
		//time.Sleep(time.Second * time.Duration(1))
	}
}

func main() {

	chn := make(chan time.Time)

	go ping(chn)
	go pong(chn)

	// just to stop with out ending
	var input string
	fmt.Scanln(&input)
}
