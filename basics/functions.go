package basics

import "fmt"

func Functions() {
	fmt.Println("Learn Go Print") // Use fmt.Println to print in new line

	/* Functions */
	/*
		//--------------------------------------------------
		// It's a function and get two argument and return sum of this two numbers
		x := Sum(2, 4)
		fmt.Printf("2 + 4 = %d", x)
		//--------------------------------------------------
	*/

}

func Sum(x int, y int) int {
	return x + y
}

func add(x, y int) int {
	return x + y
}
