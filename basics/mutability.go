package basics

import "fmt"

func Mutability() {
	fmt.Println("Learn Go Mutability") // Use fmt.Println to print in new line

	/* Pointer */
	/*
		//--------------------------------------------------
		// In Go, only constants are immutable. However, because arguments are passed by value, a function receiving a value argument and mutating it, won’t mutate the original value.
		articleV := Artist{Name: "Kasra", Genre: "Rock", Songs: 42}
		fmt.Printf("%s released their %dth song\n", articleV.Name, passByValue(articleV))
		fmt.Printf("%s has a total of %d songs", articleV.Name, articleV.Songs)

		// As you can see the total amount of songs on the article variable’s value was not changed. To mutate the passed value, we need to pass it by reference, using a pointer.
		articleR := &Artist{Name: "Kasra", Genre: "Rock", Songs: 42}
		fmt.Printf("%s released their %dth song\n", articleR.Name, passByReference(articleR))
		fmt.Printf("%s has a total of %d songs", articleR.Name, articleR.Songs)
		//--------------------------------------------------
	*/
}

type Artist struct {
	Name, Genre string
	Songs       int
}

func passByValue(a Artist) int {
	a.Songs++
	return a.Songs
}

func passByReference(a *Artist) int {
	a.Songs++
	return a.Songs
}
