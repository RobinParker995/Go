package main

import "fmt"

func main(){

	//Scanln Reads until newline
	//Good for int and float
	//It stops reading after space

	var name string

	fmt.Print("Enter your name: ");
	fmt.Scanln(&name)

	fmt.Println("Hey",name);


	fmt.Printf("\n")


	//Scan Reads space-separated input
	//Good for string inputs
	//Can read inputs for multiple variable

	var first string
	var last string

	fmt.Print("Enter your full name: ")
	fmt.Scan(&first, &last)

	fmt.Printf("First Name: %s, Last Name: %s \n", first, last)


	fmt.Printf("\n")

	//Scanf is similar to scanf from C
	//It needs format specifer for input

	var age int

	fmt.Print("Enter your age: ");
	fmt.Scanf("%d", &age);

	fmt.Printf("Age: %d \n", age);
}
