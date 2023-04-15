package syntax

import (
	"fmt"
)

func ChannelBuffering() {
	fmt.Println("Learn Go Channel Buffering")

	/* Channel Buffering */
	/*
		//--------------------------------------------------
		messages := make(chan string, 2) // Because of this channel buffered to have just 2 goroutines, if we pass more than 2 value we get fatal error: all goroutines are asleep - deadlock!

		messages <- "buffered"
		messages <- "channel"

		fmt.Println(<-messages)
		fmt.Println(<-messages)
		//--------------------------------------------------
	*/
}
