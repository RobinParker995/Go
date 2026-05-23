package main

import "fmt"

func main(){
	
	//Basic Syntax

	for i := 0; i < 5; i++ {

		fmt.Println(i)
	}

	fmt.Println()

	//Treating For as While

	x := 5

	for x > 0 {
	
		fmt.Println(x)
		x--
	}

	fmt.Println()

	//Using range for slicing

	nums := []int {10, 20, 30}

	for index, value := range nums {

		fmt.Println(index, value)
	}

	fmt.Println()


	//Use _ to ignore values
	for _, value := range nums {
 		 fmt.Println(value)
	}

}
