package __basics

import "fmt"

func VariadicFunctions() {
	fmt.Println("Learn Go VariadicFunctions")

	/* VariadicFunctions */
	/*
		//--------------------------------------------------
		VariadicSum(1)
		VariadicSum(1, 2, 3)
		VariadicSum(1, 2, 3, 4, 5, 6)
		//--------------------------------------------------
	*/
}

func VariadicSum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
