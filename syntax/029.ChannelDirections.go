package syntax

import (
	"fmt"
)

func ChannelDirections() {
	fmt.Println("Learn Go Channel Directions")

	/* Channel Directions */
	/*
		//--------------------------------------------------
		pings := make(chan string, 1)
		pongs := make(chan string, 1)
		ping(pings, "passed message")
		pong(pings, pongs)
		fmt.Println(<-pongs)
		//--------------------------------------------------
	*/
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
