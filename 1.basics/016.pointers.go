package __basics

import "fmt"

func Pointers() {
	fmt.Println("Learn Go Pointers")

	/* Pointers */
	/*
		//--------------------------------------------------
		// Explain One
		i := 1
		fmt.Println("Number is:", i)

		byValue(i)
		fmt.Println("Number after using byValue is:", i)

		byReference(&i)
		fmt.Println("Number after using byReference is:", i)

		// Explain Two Values
		first := "first value"
		sec := &first // It's something like 0x1400006cf10 because we get pointer of variable in ram, so we have that address in sec variable
		third := *sec // It's "first value" because we get value of sec variable and it's means get value of some variable address and that is address of first variable

		fmt.Println(first)
		fmt.Println(sec)
		fmt.Println(third)
		//--------------------------------------------------
	*/
}

func byValue(numberByValue int) {
	numberByValue = 0
}

func byReference(numberByReference *int) {
	*numberByReference = 0
}
