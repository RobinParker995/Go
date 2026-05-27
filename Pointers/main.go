package main

import "fmt"


//Using pointers via function to change the orignal value
func wrongDouble(s int) {
	s = s * 2 //This won't change anything as the function get copy of value
}

func double(s *int) {
	*s = *s * 2 //This will change as the function gets the address of variable which allows modification.
}

func main() {

	x := 10

	//Pointer Declration
	var p *int = &x

	//Shorthand declaration
	p2 := &x

	fmt.Println(p)
	fmt.Println(p2)

	fmt.Println()

	fmt.Println(*p) //Here * means dereference pointer which means the value stored in the address.


	fmt.Println()

	//Modifications of value using pointers
	*p2 = 76

	fmt.Println(x) //Changes the value of x

	fmt.Println()

	wrongDouble(x)
	fmt.Println(x) //Same value

	double(p2) 
	fmt.Println(x) //Doubled Value


	//We us pointers a lot in structs and methods
	//If you wanna know how and why, check out the struct and method code
	//I have written everything there
}