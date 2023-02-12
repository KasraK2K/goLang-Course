package __basics

import "fmt"

func Functions() {
	fmt.Println("Learn Go Functions")

	/* Functions */
	/*
		//--------------------------------------------------
		// It's a function and get two argument and return sum of this two numbers
		x := sum(2, 4)
		fmt.Printf("2 + 4 = %d", x)
		//--------------------------------------------------
	*/
}

func sum(x int, y int) int {
	return x + y
}

func add(x, y int) int {
	return x + y
}

func greetings(name string, age int) (string, int) {
	message := fmt.Sprintf("Hi dear %s", name)
	return message, age
}
