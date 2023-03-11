package __basics

import "fmt"

func Pointers() {
	fmt.Println("Learn Go Pointers")

	/* Pointers */
	/*
		//--------------------------------------------------
		i := 1
		fmt.Println("Number is:", i)

		byValue(i)
		fmt.Println("Number after using byValue is:", i)

		byReference(&i)
		fmt.Println("Number after using byReference is:", i)
		//--------------------------------------------------
	*/
}

func byValue(numberByValue int) {
	numberByValue = 0
}

func byReference(numberByReference *int) {
	*numberByReference = 0
}
