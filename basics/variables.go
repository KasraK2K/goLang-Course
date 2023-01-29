package basics

import "fmt"

func Variables() {
	fmt.Println("Learn Go Variables")
	/* Variables & inferred typing */
	/*
		//--------------------------------------------------
		var (
			name     string
			age      int
			location string
		)
		//--------------------------------------------------
		var (
			name, location string
			age            int
		)
		//--------------------------------------------------
		var name string
		var age int
		var location string
		//--------------------------------------------------
		var (
			name     string = "Kasra Karami"
			age      int    = 32
			location string = "Done"
		)
		//--------------------------------------------------
		var (
			name     = "Kasra Karami"
			age      = 32
			location = "Done"
		)
		//--------------------------------------------------
		var (
			name, location, age = "Kasra Karami", "Done", 32
		)
		//--------------------------------------------------
		name, location := "Kasra Karami", "Done"
		age := 32
		//--------------------------------------------------
	*/

	/* Constants */
	/*
		//--------------------------------------------------
		const Pi = 3.14
		const (
			Number = 1234
			Truth  = false
			Big    = 1 << 62
			Small  = Big >> 61
		)
		//--------------------------------------------------
	*/
}
