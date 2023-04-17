package syntax

import (
	"fmt"
)

func Select() {
	fmt.Println("Learn Go Select")

	/* Select */
	/*
		//--------------------------------------------------
		c1 := make(chan string)
		c2 := make(chan string)

		go func() {
			time.Sleep(1 * time.Millisecond)
			c1 <- "one"
		}()

		go func() {
			time.Sleep(2 * time.Second)
			c2 <- "two"
		}()

		for i := 0; i < 2; i++ {
			select {
			case msg1 := <-c1:
				fmt.Println("received", msg1)
			case msg2 := <-c2:
				fmt.Println("received", msg2)
			}
		}
		//--------------------------------------------------
	*/
}
