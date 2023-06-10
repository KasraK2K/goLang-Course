package syntax

import (
	"fmt"
)

func RangeOverChannels() {
	fmt.Println("Learn Go Range Over Channels")

	/* Range Over Channels */
	/*
		//--------------------------------------------------
		queue := make(chan string, 2)
		queue <- "one"
		queue <- "two"
		close(queue)

		for elem := range queue {
			fmt.Println(elem)
		}
		//--------------------------------------------------
	*/
}
