package basics

import (
	"fmt"
)

func Pointer() {
	fmt.Println("Learn Go Pointer") // Use fmt.Println to print in new line

	/* Pointer */
	/*
		//--------------------------------------------------
		first := "first value"
		sec := &first // It's something like 0x1400006cf10 because we get pointer of variable in ram, so we have that address in sec variable
		third := *sec // It's "first value" because we get value of sec variable and it's means get value of some variable address and that is address of first variable

		fmt.Println(first)
		fmt.Println(sec)
		fmt.Println(third)
		//--------------------------------------------------
	*/
}
