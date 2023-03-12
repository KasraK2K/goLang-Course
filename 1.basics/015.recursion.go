package __basics

import "fmt"

func Recursion() {
	fmt.Println("Learn Go Recursion")

	/* Recursion */
	/*
		//--------------------------------------------------
		// One
		fmt.Println(simpleRecursive("my text", 5))

		// Two
		fmt.Println(fact(4))

		// Three
		var fib func(n int) int
		fib = func(n int) int {
			if n < 2 {
				return n
			}
			return fib(n-1) + fib(n-2)
		}
		fmt.Println(fib(7))
		//--------------------------------------------------
	*/
}

func simpleRecursive(text string, count int) string {
	if count == 0 {
		return text
	}
	return simpleRecursive(fmt.Sprintf("%s %v", text, count), count-1)
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
