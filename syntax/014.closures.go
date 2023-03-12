package syntax

import "fmt"

func Closures() {
	fmt.Println("Learn Go Closures")

	/* Closures */
	/*
		//--------------------------------------------------
		// One
		cl1()

		// Two
		// newCounter function is
		// assigned to a variable
		counter := cl2()
		// invoke counter
		fmt.Println(counter())
		// invoke counter
		fmt.Println(counter())
		//--------------------------------------------------
	*/
}

func cl1() {
	// Declaring the variable
	GFG := 0

	// Assigning an anonymous
	// function to a variable
	counter := func() int {
		GFG += 1
		return GFG
	}

	fmt.Println(counter())
	fmt.Println(counter())
}

func cl2() func() int {
	GFG := 0
	return func() int {
		GFG += 1
		return GFG
	}
}
