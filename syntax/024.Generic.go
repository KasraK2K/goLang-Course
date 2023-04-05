package syntax

import (
	"fmt"
)

func Generic() {
	fmt.Println("Learn Go Generic")

	x := 3
	y := 5
	swap(&x, &y)
	fmt.Println(x, y)

	words := []string{"apple", "banana", "cherry", "banana", "date", "banana"}
	countOfBananas := count(words, "banana")
	fmt.Println(countOfBananas) // Output: 3

	/* Generic */
	/*
		//--------------------------------------------------
		//--------------------------------------------------
	*/
}

// example one
func swap[T any](a, b *T) {
	*a, *b = *b, *a
}

// example two
func count[T comparable](items []T, target T) int {
	count := 0
	for _, item := range items {
		if item == target {
			count++
		}
	}
	return count
}
