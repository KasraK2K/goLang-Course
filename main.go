package main

import (
	"fmt"
	"goLang-Course/1.basics"
	"goLang-Course/2.types"
)

func main() {
	fmt.Println("goLang Tutorial with code and comment") // Example of print message

	// Basics
	__basics.Variables() // In this line `basics` is name of package we wrote in variables file eq: package basics
	__basics.Constants()
	__basics.Print()
	__basics.Functions()
	__basics.Pointer()
	__basics.For()
	__basics.Mutability()
	__basics.If()

	// Types
	__types.BasicTypes()
	__types.TypeConversion()
}
