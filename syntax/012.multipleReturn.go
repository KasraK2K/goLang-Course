package syntax

import "fmt"

func MultipleReturn() {
	fmt.Println("Learn Go MultipleReturn")

	/* MultipleReturn */
	/*
		//--------------------------------------------------
		c, d := myReturnFunction()
		fmt.Println(c, d)
		//--------------------------------------------------
	*/
}

func myReturnFunction() (int, int) {
	return 2, 4
}
