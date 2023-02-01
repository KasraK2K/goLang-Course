package main

import (
	/**
	Import `fmt` package to use it for print a message in console
	*/
	"fmt"
	/**
	`goLang-Course` is our golang application name and `basics` is our directory name.
	In this package should declare our module name and that is `basics`. eq: package basics
	Into this directory we have many modules like `variables.go` file. So we should use `<package_name>.` before function name. eq: basics.Variables()
	*/
	"goLang-Course/basics"
)

func main() {
	fmt.Println("goLang Tutorial with code and comment") // Example of print message

	basics.Variables() // In this line `basics` is name of package we wrote in variables file eq: package basics
	basics.Constants()
	basics.Print()
	basics.Functions()
	basics.Pointer()
	basics.For()
}
