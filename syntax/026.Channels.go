package syntax

import (
	"fmt"
)

func Channels() {
	fmt.Println("Learn Go Channels")

	/* Channels */
	/*
		//--------------------------------------------------
		messages := make(chan string)

		go func() { messages <- "ping" }()

		msg := <-messages
		fmt.Println(msg)
		//--------------------------------------------------
	*/
}
