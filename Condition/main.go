package main

import "fmt"

func main(){

	var age int

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	if age<=17 {

		fmt.Println("You can't apparate")

	} else {

		fmt.Println("You can apparate")
	}
}
