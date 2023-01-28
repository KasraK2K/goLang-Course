package main

import "fmt"                     // Import `fmt` package to use it for print a message in console
import "goLang-Course/variables" // `goLang-Course` is our golang project name and `variables` is our directory name. Into this directory we have `variables.go` file so we should use `variables.` before function name.

func main() {
	fmt.Println("goLang Tutorial with code and comment") // Example of print message

	variables.Greeting() // In this line `variables` in name of package we wrote in variables file eq: package variables
}
