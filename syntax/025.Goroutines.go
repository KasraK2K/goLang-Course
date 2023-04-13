package syntax

import (
	"fmt"
)

func Goroutines() {
	fmt.Println("Learn Go Goroutines")

	/* Goroutines */
	/*
		//--------------------------------------------------
		f("direct") // this function runs synchronously

		go f("goroutine") // this function runs asynchronously
	
		go func(msg string) {
			fmt.Println(msg)
		}("going") // this function runs asynchronously

		time.Sleep(time.Second) // When we run function on goroutine, we need to wait for goroutine to finish
		fmt.Println("done")
		//--------------------------------------------------
	*/
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
