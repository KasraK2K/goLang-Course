package main

import (
	"fmt"
	"goLang-Course/1.basics"
	__types "goLang-Course/2.types"
)

func main() {
	fmt.Println("goLang Tutorial with code and comment") // Example of print message

	// Basics
	__basics.Variables() // In this line `basics` is name of package we wrote in variables file eq: package basics
	__basics.Constants()
	__basics.Print()
	__basics.Pointer()
	__basics.For()
	__basics.Mutability()
	__basics.If()
	__basics.Switch()
	__basics.Array()
	__basics.Slice()
	__basics.Map()
	__basics.Range()
	__basics.Functions()
	__basics.MultipleReturn()
	__basics.VariadicFunctions()

	// Types
	__types.BasicTypes()
	__types.TypeConversion()
}
