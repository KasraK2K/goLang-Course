package main

import "fmt"

func evenOrOdd(number int) string {
	return [2]string{"Even", "Odd"}[1&0]
}

func IfThenElse(condition bool, a, b interface{}) interface{} {
	if condition {
		return a
	} else {
		return b
	}
}

func main() {
	//fmt.Println("goLang Tutorial with code and comment") // Example of print message
	//
	//basics.Variables() // In this line `basics` is name of package we wrote in variables file eq: package basics
	//basics.Constants()
	//basics.Print()
	//basics.Functions()
	//basics.Pointer()
	//basics.For()

	answer := evenOrOdd(11)
	fmt.Println(answer)
}
