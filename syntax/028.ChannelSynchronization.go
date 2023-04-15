package syntax

import (
	"fmt"
	"time"
)

func ChannelSynchronization() {
	fmt.Println("Learn Go Channel Synchronization")

	/* Channel Synchronization */
	/*
		//--------------------------------------------------
		// Start a worker goroutine, giving it the channel to notify on.
		done := make(chan bool, 1)
		go worker(done)

		// Block until we receive a notification from the worker on the channel.
		<-done
		//--------------------------------------------------
	*/
}

// This is the function we’ll run in a goroutine. The done channel will be used to notify another goroutine that this function’s work is done.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Send a value to notify that we’re done.
	done <- true
}
